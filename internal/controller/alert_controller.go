/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	monitoringv1alpha1 "github.com/harsh098/signoz-alert-operator/api/v1alpha1"
	"github.com/harsh098/signoz-alert-operator/internal/signozclient"
)

const retryInterval = 5

// AlertReconciler reconciles a Alert object
type AlertReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const finalizerName = "monitoring.hmx86.cloud/finalizer"

// +kubebuilder:rbac:groups=monitoring.hmx86.cloud,resources=alerts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitoring.hmx86.cloud,resources=alerts/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitoring.hmx86.cloud,resources=alerts/finalizers,verbs=update
// +kubebuilder:rbac:groups=monitoring.hmx86.cloud,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Alert object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *AlertReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var alert monitoringv1alpha1.Alert
	var apiClient *signozclient.ClientWithResponses
	var err error
	logger := logf.FromContext(ctx)
	if err := r.Get(ctx, req.NamespacedName, &alert); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	oldStatus := alert.Status
	hasFinaliser := controllerutil.ContainsFinalizer(&alert, finalizerName)

	if alert.DeletionTimestamp == nil && !hasFinaliser {
		controllerutil.AddFinalizer(&alert, finalizerName)
		if err := r.Update(ctx, &alert); err != nil {
			return ctrl.Result{}, err
		}
	}

	apiClient, err = r.resolveClient(ctx, &alert)

	if err != nil {
		alert.Status.Errors = err.Error()
		logger.Error(err, "Failed to resolve Signoz Client")
		_ = r.Status().Update(ctx, &alert)
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	var foundByID bool
	var adoptableID string
	var httpStatus int
	var errBody string
	var op alertOp

	foundByID, adoptableID, httpStatus, errBody, err = r.probe(ctx, apiClient, &alert)
	alert.Status.Errors = errBody
	alert.Status.HTTPStatus = httpStatus
	if err != nil {
		alert.Status.Errors = fmt.Sprintf("Error: %v; errBody: %v", err.Error(), errBody)
		logger.Error(err, "Failed to Query Signoz API")
		_ = r.Status().Update(ctx, &alert)
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	if httpStatus >= 400 {
		logger.Error(errors.New("signoz returned non-2xx during probe"), "scheduling retry", "status", httpStatus, "body", errBody)
		_ = r.Status().Update(ctx, &alert)
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	op = plan(&alert, foundByID, adoptableID)

	httpStatus, errBody, err = r.apply(ctx, apiClient, &alert, op, adoptableID)
	alert.Status.HTTPStatus = httpStatus
	alert.Status.Errors = errBody

	if err != nil {
		alert.Status.Errors = fmt.Sprintf("Error: %v; errBody: %v", err.Error(), errBody)
		logger.Error(err, "Failed to Query Signoz API")
		_ = r.Status().Update(ctx, &alert)
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	if httpStatus >= 400 {
		logger.Error(errors.New("signoz returned non-2xx"), "scheduling retry", "status", httpStatus, "body", errBody)
		_ = r.Status().Update(ctx, &alert)
		isCleanup := httpStatus == 404 && op == DELET
		if !isCleanup {
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}
	}

	if alert.DeletionTimestamp != nil {
		controllerutil.RemoveFinalizer(&alert, finalizerName)
		if err := r.Update(ctx, &alert); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	if !reflect.DeepEqual(oldStatus, alert.Status) {
		_ = r.Status().Update(ctx, &alert)
	}

	requeInterval := wait.Jitter(retryInterval*time.Minute, 0.1)
	return ctrl.Result{RequeueAfter: requeInterval}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AlertReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringv1alpha1.Alert{}).
		Named("alert").
		Complete(r)
}

// Helpers used by Reconcile.

// k8sID is the idempotency key stamped on every SigNoz rule's `labels.k8s_id`.
// Same Alert manifest applied in two clusters that share a SigNoz instance
// produces the same k8s_id, enabling cross-cluster manifest migration without
// duplicating rules.
func k8sID(alert *monitoringv1alpha1.Alert) string {
	return alert.Namespace + "-" + alert.Name
}

// resolveClient fetches the Endpoint CR referenced by the Alert, reads the API
// key from the referenced Secret, and returns a SigNoz client that injects the
// SigNoz-Api-Key header on every request.
func (r *AlertReconciler) resolveClient(ctx context.Context, alert *monitoringv1alpha1.Alert) (*signozclient.ClientWithResponses, error) {
	var endpoint monitoringv1alpha1.Endpoint
	if err := r.Get(ctx, types.NamespacedName{
		Name:      alert.Spec.EndpointRef.Name,
		Namespace: alert.Namespace,
	}, &endpoint); err != nil {
		return nil, fmt.Errorf("fetch endpoint %q: %w", alert.Spec.EndpointRef.Name, err)
	}

	if endpoint.Spec.SecretKeyRef == nil {
		return nil, fmt.Errorf("endpoint %q has no secretKeyRef set", endpoint.Name)
	}

	var secret corev1.Secret
	if err := r.Get(ctx, types.NamespacedName{
		Name:      endpoint.Spec.SecretKeyRef.Name,
		Namespace: alert.Namespace,
	}, &secret); err != nil {
		return nil, fmt.Errorf("fetch secret %q: %w", endpoint.Spec.SecretKeyRef.Name, err)
	}
	keyBytes, ok := secret.Data[endpoint.Spec.SecretKeyRef.Key]
	if !ok {
		return nil, fmt.Errorf("secret %q has no key %q", secret.Name, endpoint.Spec.SecretKeyRef.Key)
	}
	apiKey := string(keyBytes)

	return signozclient.NewClientWithResponses(
		endpoint.Spec.InstanceURL,
		signozclient.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("SigNoz-Api-Key", apiKey)
			return nil
		}),
	)
}

// probe asks SigNoz what it currently knows about this Alert. Two-step:
//  1. If status.RuleID is set, GET /api/v2/rules/{id}. 200 → foundByID = true.
//     404 → fall through to list-and-match (rule was deleted via UI). Other →
//     surface the http status/body, caller decides what to do.
//  2. List /api/v2/rules and look for one tagged with labels.k8s_id == k8sID(alert).
//     Match → adoptableID is its SigNoz id. No match → both bools/strings zero.
//
// Returns (foundByID, adoptableID, httpStatus, errBody, err). err is for
// transport-level failures only; HTTP errors are surfaced via httpStatus/errBody.
func (r *AlertReconciler) probe(ctx context.Context, api *signozclient.ClientWithResponses, alert *monitoringv1alpha1.Alert) (foundByID bool, adoptableID string, httpStatus int, errBody string, err error) {
	if alert.Status.RuleID != "" {
		resp, gerr := api.GetRuleByIDWithResponse(ctx, alert.Status.RuleID)
		if gerr != nil {
			return false, "", 0, "", fmt.Errorf("get rule %q: %w", alert.Status.RuleID, gerr)
		}
		switch resp.StatusCode() {
		case http.StatusOK:
			return true, "", http.StatusOK, "", nil
		case http.StatusNotFound:
			// Manually deleted via SigNoz UI. Fall through to list-and-match.
		default:
			return false, "", resp.StatusCode(), string(resp.Body), nil
		}
	}

	listResp, lerr := api.ListRulesWithResponse(ctx)
	if lerr != nil {
		return false, "", 0, "", fmt.Errorf("list rules: %w", lerr)
	}
	if listResp.JSON200 == nil {
		return false, "", listResp.StatusCode(), string(listResp.Body), nil
	}
	want := k8sID(alert)
	for _, rule := range listResp.JSON200.Data {
		if rule.Labels != nil && (*rule.Labels)["k8s_id"] == want {
			return false, rule.Id, http.StatusOK, "", nil
		}
	}
	return false, "", http.StatusOK, "", nil
}

// apply executes the planned operation against SigNoz.
//
// For ADOPT, apply writes adoptableID into alert.Status.RuleID before PUTting —
// this is the cross-cluster recovery path: probe found an existing rule tagged
// with our k8s_id, and we claim it by setting our status.RuleID to its id,
// then fall through to the same PUT logic as UPDAT.
//
// On successful CREAT, apply writes the new rule id back into
// alert.Status.RuleID in-memory; the actual K8s persist happens later in
// Reconcile via Status().Update.
func (r *AlertReconciler) apply(ctx context.Context, api *signozclient.ClientWithResponses, alert *monitoringv1alpha1.Alert, op alertOp, adoptableID string) (httpStatus int, errBody string, err error) {
	switch op {
	case NOP:
		return 0, "", nil

	case CREAT:
		body, berr := buildRuleBody(alert)
		if berr != nil {
			return 0, "", berr
		}
		resp, perr := api.CreateRuleWithResponse(ctx, body)
		if perr != nil {
			return 0, "", fmt.Errorf("create rule: %w", perr)
		}
		if resp.JSON201 != nil {
			alert.Status.RuleID = resp.JSON201.Data.Id
			return http.StatusCreated, "", nil
		}
		return resp.StatusCode(), string(resp.Body), nil

	case ADOPT:
		// Claim the existing rule by stamping its id into our status, then
		// fall through to the standard PUT path.
		alert.Status.RuleID = adoptableID
		fallthrough

	case UPDAT:
		body, berr := buildRuleBody(alert)
		if berr != nil {
			return 0, "", berr
		}
		resp, perr := api.UpdateRuleByIDWithResponse(ctx, alert.Status.RuleID, body)
		if perr != nil {
			return 0, "", fmt.Errorf("update rule %q: %w", alert.Status.RuleID, perr)
		}
		if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
			return resp.StatusCode(), "", nil
		}
		return resp.StatusCode(), string(resp.Body), nil

	case DELET:
		resp, perr := api.DeleteRuleByIDWithResponse(ctx, alert.Status.RuleID)
		if perr != nil {
			return 0, "", fmt.Errorf("delete rule %q: %w", alert.Status.RuleID, perr)
		}
		// 204 = deleted, 404 = already gone — both are success for cleanup.
		if resp.StatusCode() == http.StatusNoContent || resp.StatusCode() == http.StatusNotFound {
			return resp.StatusCode(), "", nil
		}
		return resp.StatusCode(), string(resp.Body), nil
	}
	return 0, "", nil
}

// buildRuleBody decodes the Alert's RawExtension spec.rule into the typed
// SigNoz body and stamps the k8s_id idempotency label.
func buildRuleBody(alert *monitoringv1alpha1.Alert) (signozclient.RuletypesPostableRule, error) {
	var body signozclient.RuletypesPostableRule
	if err := json.Unmarshal(alert.Spec.Rule.Raw, &body); err != nil {
		return body, fmt.Errorf("decode spec.rule: %w", err)
	}
	if body.Labels == nil {
		m := map[string]string{}
		body.Labels = &m
	}
	(*body.Labels)["k8s_id"] = k8sID(alert)
	return body, nil
}
