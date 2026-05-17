/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	api "github.com/harsh098/signoz-alert-operator/api/v1alpha1"
)

const (
	testAPIKey   = "test-api-key"
	rulesPath    = "/api/v2/rules"
	testRuleID   = "rule-42"
	testRulePath = rulesPath + "/" + testRuleID
)

// signozCall captures one request the controller made against the fake SigNoz.
type signozCall struct {
	method string
	path   string
	apiKey string
	body   []byte
}

// fakeSignoz is a tiny httptest harness that records calls and lets each spec
// override the handler. Default handler returns 200 with empty bodies so a
// misconfigured test fails loudly instead of silently passing.
type fakeSignoz struct {
	mu      sync.Mutex
	server  *httptest.Server
	calls   []signozCall
	handler http.HandlerFunc
}

func newFakeSignoz() *fakeSignoz {
	f := &fakeSignoz{}
	f.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		f.calls = append(f.calls, signozCall{
			method: r.Method,
			path:   r.URL.Path,
			apiKey: r.Header.Get("SigNoz-Api-Key"),
			body:   body,
		})
		h := f.handler
		f.mu.Unlock()
		h(w, r)
	}))
	return f
}

func (f *fakeSignoz) setHandler(h http.HandlerFunc) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.handler = h
}

func (f *fakeSignoz) recorded() []signozCall {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]signozCall, len(f.calls))
	copy(out, f.calls)
	return out
}

func (f *fakeSignoz) close() { f.server.Close() }

var _ = Describe("Alert Controller", func() {
	var (
		ctx        context.Context
		fakeAPI    *fakeSignoz
		reconciler *AlertReconciler
		secret     *corev1.Secret
		endpoint   *api.Endpoint
		alert      *api.Alert
		nsName     types.NamespacedName
	)

	BeforeEach(func() {
		ctx = context.Background()
		fakeAPI = newFakeSignoz()

		// Secret holding the API key.
		secret = &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{Name: "signoz-creds", Namespace: "default"},
			StringData: map[string]string{"api-key": testAPIKey},
		}
		Expect(k8sClient.Create(ctx, secret)).To(Succeed())

		// Endpoint pointing at the fake SigNoz.
		endpoint = &api.Endpoint{
			ObjectMeta: metav1.ObjectMeta{Name: "signoz", Namespace: "default"},
			Spec: api.EndpointSpec{
				InstanceURL:  fakeAPI.server.URL,
				SecretKeyRef: &api.SecretKeyRefSpec{Name: secret.Name, Key: "api-key"},
			},
		}
		Expect(k8sClient.Create(ctx, endpoint)).To(Succeed())

		// Alert via fixture. Namespace overridden to "default" for envtest.
		alert = newAlert(withNamespace("default"), withEndpoint(endpoint.Name))
		Expect(k8sClient.Create(ctx, alert)).To(Succeed())

		nsName = types.NamespacedName{Name: alert.Name, Namespace: alert.Namespace}
		reconciler = &AlertReconciler{Client: k8sClient, Scheme: k8sClient.Scheme()}
	})

	AfterEach(func() {
		// Strip any finalizer Reconcile may have added — otherwise K8s leaves the
		// object in "marked for deletion" state across tests and the next BeforeEach
		// 409s with AlreadyExists. The real controller's DELET path normally cleans
		// finalizers, but no controller runs between specs in envtest.
		fresh := &api.Alert{}
		if err := k8sClient.Get(ctx, nsName, fresh); err == nil {
			fresh.Finalizers = nil
			_ = k8sClient.Update(ctx, fresh)
			_ = k8sClient.Delete(ctx, fresh)
		}
		_ = client.IgnoreNotFound(k8sClient.Delete(ctx, endpoint))
		_ = client.IgnoreNotFound(k8sClient.Delete(ctx, secret))
		fakeAPI.close()
	})

	Context("on a fresh Alert (no status.ruleID)", func() {
		It("creates the rule in SigNoz, stamps k8s_id, and persists ruleID + finalizer", func() {
			// Fake SigNoz: GET /rules (list, used by probe) → empty, POST /rules → 201 with id.
			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == rulesPath:
					_ = json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": []any{}})
				case r.Method == http.MethodPost && r.URL.Path == rulesPath:
					w.WriteHeader(http.StatusCreated)
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": testRuleID, "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})
			Expect(err).NotTo(HaveOccurred())

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("populating status from the SigNoz response")
			Expect(updated.Status.RuleID).To(Equal(testRuleID))
			Expect(updated.Status.HTTPStatus).To(Equal(http.StatusCreated))
			Expect(updated.Status.Errors).To(BeEmpty())

			By("adding the finalizer")
			Expect(updated.Finalizers).To(ContainElement(finalizerName))

			By("authenticating with the SigNoz-Api-Key header")
			calls := fakeAPI.recorded()
			Expect(calls).NotTo(BeEmpty())
			for _, c := range calls {
				Expect(c.apiKey).To(Equal(testAPIKey))
			}

			By("issuing exactly one POST /api/v2/rules")
			posts := filterCalls(calls, http.MethodPost, rulesPath)
			Expect(posts).To(HaveLen(1))

			By("stamping k8s_id=<ns>-<name> on labels in the POST body")
			var sent map[string]any
			Expect(json.Unmarshal(posts[0].body, &sent)).To(Succeed())
			labels, _ := sent["labels"].(map[string]any)
			Expect(labels).NotTo(BeNil(), "POST body missing labels")
			Expect(labels["k8s_id"]).To(Equal("default-" + alert.Name))
		})
	})

	// ----- Pending specs below: enable as you implement each path -----

	Context("when status.ruleID exists and the rule is present in SigNoz", func() {
		It("issues PUT /api/v2/rules/{id} (no POST, no list-all)", func() {
			// Pre-seed status.ruleID — Alert.Status is a subresource, so use Status().Update.
			alert.Status.RuleID = testRuleID
			Expect(k8sClient.Status().Update(ctx, alert)).To(Succeed())

			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == testRulePath:
					// Probe: rule exists.
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": testRuleID, "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				case r.Method == http.MethodPut && r.URL.Path == testRulePath:
					// Update succeeded.
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": testRuleID, "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})
			Expect(err).NotTo(HaveOccurred())

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("keeping the same ruleID")
			Expect(updated.Status.RuleID).To(Equal(testRuleID))
			Expect(updated.Status.HTTPStatus).To(Equal(http.StatusOK))
			Expect(updated.Status.Errors).To(BeEmpty())

			By("keeping the finalizer")
			Expect(updated.Finalizers).To(ContainElement(finalizerName))

			calls := fakeAPI.recorded()

			By("issuing zero POSTs (no duplicate creates)")
			Expect(filterCalls(calls, http.MethodPost, rulesPath)).To(BeEmpty())

			By("issuing zero list-all GETs (steady-state must not scan all rules)")
			Expect(filterCalls(calls, http.MethodGet, rulesPath)).To(BeEmpty())

			By("issuing exactly one PUT /api/v2/rules/rule-42")
			Expect(filterCalls(calls, http.MethodPut, testRulePath)).To(HaveLen(1))

			By("authenticating every call with SigNoz-Api-Key")
			for _, c := range calls {
				Expect(c.apiKey).To(Equal(testAPIKey))
			}
		})
	})

	Context("when status.ruleID exists but the rule was deleted from the SigNoz UI", func() {
		It("falls through to POST and updates status.ruleID with the new id", func() {
			// Pre-seed a stale ruleID — simulates "successfully reconciled
			// rule-42 earlier, then a human deleted it via the SigNoz UI."
			alert.Status.RuleID = testRuleID
			Expect(k8sClient.Status().Update(ctx, alert)).To(Succeed())

			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == testRulePath:
					// Probe by id: rule is gone. SigNoz returns a JSON error body
					// here (not an empty 404), so we mirror that — otherwise
					// oapi-codegen's parser errors on "unexpected end of JSON input".
					w.WriteHeader(http.StatusNotFound)
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "error",
						"error": map[string]any{
							"code":    "not_found",
							"message": "rule with ID: rule-42 does not exist",
						},
					})
				case r.Method == http.MethodGet && r.URL.Path == rulesPath:
					// Probe falls through to list — and nothing matches our k8s_id.
					_ = json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": []any{}})
				case r.Method == http.MethodPost && r.URL.Path == rulesPath:
					w.WriteHeader(http.StatusCreated)
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": "rule-99", "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})
			Expect(err).NotTo(HaveOccurred())

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("replacing the stale ruleID with the newly created one")
			Expect(updated.Status.RuleID).To(Equal("rule-99"))
			Expect(updated.Status.HTTPStatus).To(Equal(http.StatusCreated))
			Expect(updated.Status.Errors).To(BeEmpty())

			calls := fakeAPI.recorded()

			By("probing by stale id, then listing, then creating")
			Expect(filterCalls(calls, http.MethodGet, testRulePath)).To(HaveLen(1))
			Expect(filterCalls(calls, http.MethodGet, rulesPath)).To(HaveLen(1))
			Expect(filterCalls(calls, http.MethodPost, rulesPath)).To(HaveLen(1))

			By("not issuing any PUT (the old rule was gone, the new one was just created)")
			Expect(filterCalls(calls, http.MethodPut, testRulePath)).To(BeEmpty())
			Expect(filterCalls(calls, http.MethodPut, "/api/v2/rules/rule-99")).To(BeEmpty())

			By("authenticating every call with SigNoz-Api-Key")
			for _, c := range calls {
				Expect(c.apiKey).To(Equal(testAPIKey))
			}
		})
	})

	Context("when status.ruleID is empty but an adoptable rule exists in SigNoz", func() {
		It("adopts the existing rule (PUT, not POST)", func() {
			// BeforeEach already created the Alert with empty status.RuleID —
			// the fresh, never-reconciled-in-this-cluster state. Simulates:
			// "another cluster's operator created this rule and labelled it
			// with our k8s_id; we should claim it via PUT, not create a duplicate."

			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == rulesPath:
					// List returns one rule already tagged with our k8s_id.
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data": []map[string]any{
							{
								"id":        "rule-77",
								"alert":     "test-alert",
								"alertType": "METRIC_BASED_ALERT",
								"ruleType":  "threshold_rule",
								"condition": map[string]any{"compositeQuery": map[string]any{}},
								"state":     "inactive",
								"labels":    map[string]string{"k8s_id": "default-" + alert.Name},
							},
						},
					})
				case r.Method == http.MethodPut && r.URL.Path == "/api/v2/rules/rule-77":
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": "rule-77", "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})
			Expect(err).NotTo(HaveOccurred())

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("adopting the existing rule's id into status.RuleID")
			Expect(updated.Status.RuleID).To(Equal("rule-77"))
			Expect(updated.Status.HTTPStatus).To(Equal(http.StatusOK))
			Expect(updated.Status.Errors).To(BeEmpty())

			calls := fakeAPI.recorded()

			By("listing first, then PUTting to the adopted id")
			Expect(filterCalls(calls, http.MethodGet, rulesPath)).To(HaveLen(1))
			Expect(filterCalls(calls, http.MethodPut, "/api/v2/rules/rule-77")).To(HaveLen(1))

			By("not issuing any POST (no duplicate rule)")
			Expect(filterCalls(calls, http.MethodPost, rulesPath)).To(BeEmpty())

			By("authenticating every call with SigNoz-Api-Key")
			for _, c := range calls {
				Expect(c.apiKey).To(Equal(testAPIKey))
			}
		})
	})

	Context("when the Alert is being deleted with a known ruleID", func() {
		It("calls DELETE on SigNoz then removes the finalizer", func() {
			// Pre-seed: install the finalizer + a ruleID, simulating an Alert that
			// has been successfully reconciled at least once.
			alert.Finalizers = append(alert.Finalizers, finalizerName)
			Expect(k8sClient.Update(ctx, alert)).To(Succeed())
			alert.Status.RuleID = testRuleID
			Expect(k8sClient.Status().Update(ctx, alert)).To(Succeed())

			// Mark the Alert for deletion. With the finalizer present, K8s sets
			// DeletionTimestamp instead of immediately GCing the object.
			Expect(k8sClient.Delete(ctx, alert)).To(Succeed())

			// Sanity: object still around with DeletionTimestamp set.
			marked := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, marked)).To(Succeed())
			Expect(marked.DeletionTimestamp).NotTo(BeNil())

			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == testRulePath:
					// Probe: rule still exists.
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "success",
						"data":   map[string]any{"id": testRuleID, "alert": "test-alert", "alertType": "METRIC_BASED_ALERT", "ruleType": "threshold_rule", "condition": map[string]any{"compositeQuery": map[string]any{}}, "state": "inactive"},
					})
				case r.Method == http.MethodDelete && r.URL.Path == testRulePath:
					w.WriteHeader(http.StatusNoContent)
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})
			Expect(err).NotTo(HaveOccurred())

			calls := fakeAPI.recorded()

			By("issuing exactly one DELETE /api/v2/rules/rule-42")
			Expect(filterCalls(calls, http.MethodDelete, testRulePath)).To(HaveLen(1))

			By("not issuing any POST or PUT (deletion path only)")
			Expect(filterCalls(calls, http.MethodPost, rulesPath)).To(BeEmpty())
			Expect(filterCalls(calls, http.MethodPut, testRulePath)).To(BeEmpty())

			By("authenticating every call with SigNoz-Api-Key")
			for _, c := range calls {
				Expect(c.apiKey).To(Equal(testAPIKey))
			}

			By("removing the finalizer so K8s GCs the Alert")
			gone := &api.Alert{}
			getErr := k8sClient.Get(ctx, nsName, gone)
			Expect(errors.IsNotFound(getErr)).To(BeTrue(), "expected Alert to be GC'd; got err: %v", getErr)
		})
	})

	Context("when SigNoz returns a non-2xx response", func() {
		It("records HTTPStatus + Errors on status and signals a retry", func() {
			fakeAPI.setHandler(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodGet && r.URL.Path == rulesPath:
					_ = json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": []any{}})
				case r.Method == http.MethodPost && r.URL.Path == rulesPath:
					w.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(w).Encode(map[string]any{
						"status": "error",
						"error":  map[string]any{"code": "internal_error", "message": "database is sad"},
					})
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			})

			result, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})

			By("signaling a retry (err returned OR requeue scheduled)")
			retryScheduled := err != nil || result.RequeueAfter > 0
			Expect(retryScheduled).To(BeTrue(), "expected reconcile to schedule a retry for 5xx; got result=%+v err=%v", result, err)

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("recording the HTTP status from SigNoz on Status.HTTPStatus")
			Expect(updated.Status.HTTPStatus).To(Equal(http.StatusInternalServerError))

			By("recording the error body on Status.Errors")
			Expect(updated.Status.Errors).NotTo(BeEmpty())
			Expect(updated.Status.Errors).To(ContainSubstring("database is sad"))

			By("not setting RuleID (the POST failed)")
			Expect(updated.Status.RuleID).To(BeEmpty())
		})
	})

	Context("when the referenced Endpoint is missing", func() {
		It("records the error on status and does NOT call SigNoz", func() {
			// Delete the Endpoint that BeforeEach created so resolveClient fails.
			Expect(k8sClient.Delete(ctx, endpoint)).To(Succeed())

			result, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: nsName})

			By("signaling a retry (err returned OR requeue scheduled)")
			retryScheduled := err != nil || result.RequeueAfter > 0
			Expect(retryScheduled).To(BeTrue(), "expected reconcile to schedule a retry; got result=%+v err=%v", result, err)

			updated := &api.Alert{}
			Expect(k8sClient.Get(ctx, nsName, updated)).To(Succeed())

			By("recording the resolve failure on Status.Errors")
			Expect(updated.Status.Errors).NotTo(BeEmpty())
			Expect(updated.Status.Errors).To(ContainSubstring("endpoint"))

			By("not making any SigNoz calls")
			Expect(fakeAPI.recorded()).To(BeEmpty())
		})
	})
})

func filterCalls(calls []signozCall, method, path string) []signozCall {
	var out []signozCall
	for _, c := range calls {
		if c.method == method && c.path == path {
			out = append(out, c)
		}
	}
	return out
}

// Suppress "imported and not used" if the suite ever drops these in refactors.
var _ = errors.IsNotFound
