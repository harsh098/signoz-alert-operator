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

package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/harsh098/signoz-alert-operator/test/utils"
)

// namespace where the project is deployed in
const namespace = "signoz-alert-operator-system"

// serviceAccountName created for the project
const serviceAccountName = "signoz-alert-operator-controller-manager"

// metricsServiceName is the name of the metrics service of the project
const metricsServiceName = "signoz-alert-operator-controller-manager-metrics-service"

// metricsRoleBindingName is the name of the RBAC that will be created to allow get the metrics data
const metricsRoleBindingName = "signoz-alert-operator-metrics-binding"

var _ = Describe("Manager", Ordered, func() {
	var controllerPodName string

	// Before running the tests, set up the environment by creating the namespace,
	// enforce the restricted security policy to the namespace, installing CRDs,
	// and deploying the controller.
	BeforeAll(func() {
		By("creating manager namespace")
		cmd := exec.Command("kubectl", "create", "ns", namespace)
		_, err := utils.Run(cmd)
		Expect(err).NotTo(HaveOccurred(), "Failed to create namespace")

		By("labeling the namespace to enforce the restricted security policy")
		cmd = exec.Command("kubectl", "label", "--overwrite", "ns", namespace,
			"pod-security.kubernetes.io/enforce=restricted")
		_, err = utils.Run(cmd)
		Expect(err).NotTo(HaveOccurred(), "Failed to label namespace with restricted policy")

		By("installing CRDs")
		cmd = exec.Command("make", "install")
		_, err = utils.Run(cmd)
		Expect(err).NotTo(HaveOccurred(), "Failed to install CRDs")

		By("deploying the controller-manager")
		cmd = exec.Command("make", "deploy", fmt.Sprintf("IMG=%s", projectImage))
		_, err = utils.Run(cmd)
		Expect(err).NotTo(HaveOccurred(), "Failed to deploy the controller-manager")
	})

	// After all tests have been executed, clean up by undeploying the controller, uninstalling CRDs,
	// and deleting the namespace.
	AfterAll(func() {
		By("cleaning up the curl pod for metrics")
		cmd := exec.Command("kubectl", "delete", "pod", "curl-metrics", "-n", namespace)
		_, _ = utils.Run(cmd)

		By("undeploying the controller-manager")
		cmd = exec.Command("make", "undeploy")
		_, _ = utils.Run(cmd)

		By("uninstalling CRDs")
		cmd = exec.Command("make", "uninstall")
		_, _ = utils.Run(cmd)

		By("removing manager namespace")
		cmd = exec.Command("kubectl", "delete", "ns", namespace)
		_, _ = utils.Run(cmd)
	})

	// After each test, check for failures and collect logs, events,
	// and pod descriptions for debugging.
	AfterEach(func() {
		specReport := CurrentSpecReport()
		if specReport.Failed() {
			By("Fetching controller manager pod logs")
			cmd := exec.Command("kubectl", "logs", controllerPodName, "-n", namespace)
			controllerLogs, err := utils.Run(cmd)
			if err == nil {
				_, _ = fmt.Fprintf(GinkgoWriter, "Controller logs:\n %s", controllerLogs)
			} else {
				_, _ = fmt.Fprintf(GinkgoWriter, "Failed to get Controller logs: %s", err)
			}

			By("Fetching Kubernetes events")
			cmd = exec.Command("kubectl", "get", "events", "-n", namespace, "--sort-by=.lastTimestamp")
			eventsOutput, err := utils.Run(cmd)
			if err == nil {
				_, _ = fmt.Fprintf(GinkgoWriter, "Kubernetes events:\n%s", eventsOutput)
			} else {
				_, _ = fmt.Fprintf(GinkgoWriter, "Failed to get Kubernetes events: %s", err)
			}

			By("Fetching curl-metrics logs")
			cmd = exec.Command("kubectl", "logs", "curl-metrics", "-n", namespace)
			metricsOutput, err := utils.Run(cmd)
			if err == nil {
				_, _ = fmt.Fprintf(GinkgoWriter, "Metrics logs:\n %s", metricsOutput)
			} else {
				_, _ = fmt.Fprintf(GinkgoWriter, "Failed to get curl-metrics logs: %s", err)
			}

			By("Fetching controller manager pod description")
			cmd = exec.Command("kubectl", "describe", "pod", controllerPodName, "-n", namespace)
			podDescription, err := utils.Run(cmd)
			if err == nil {
				fmt.Println("Pod description:\n", podDescription)
			} else {
				fmt.Println("Failed to describe controller pod")
			}
		}
	})

	SetDefaultEventuallyTimeout(2 * time.Minute)
	SetDefaultEventuallyPollingInterval(time.Second)

	Context("Manager", func() {
		It("should run successfully", func() {
			By("validating that the controller-manager pod is running as expected")
			verifyControllerUp := func(g Gomega) {
				// Get the name of the controller-manager pod
				cmd := exec.Command("kubectl", "get",
					"pods", "-l", "control-plane=controller-manager",
					"-o", "go-template={{ range .items }}"+
						"{{ if not .metadata.deletionTimestamp }}"+
						"{{ .metadata.name }}"+
						"{{ \"\\n\" }}{{ end }}{{ end }}",
					"-n", namespace,
				)

				podOutput, err := utils.Run(cmd)
				g.Expect(err).NotTo(HaveOccurred(), "Failed to retrieve controller-manager pod information")
				podNames := utils.GetNonEmptyLines(podOutput)
				g.Expect(podNames).To(HaveLen(1), "expected 1 controller pod running")
				controllerPodName = podNames[0]
				g.Expect(controllerPodName).To(ContainSubstring("controller-manager"))

				// Validate the pod's status
				cmd = exec.Command("kubectl", "get",
					"pods", controllerPodName, "-o", "jsonpath={.status.phase}",
					"-n", namespace,
				)
				output, err := utils.Run(cmd)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(output).To(Equal("Running"), "Incorrect controller-manager pod status")
			}
			Eventually(verifyControllerUp).Should(Succeed())
		})

		It("should ensure the metrics endpoint is serving metrics", func() {
			By("creating a ClusterRoleBinding for the service account to allow access to metrics")
			cmd := exec.Command("kubectl", "create", "clusterrolebinding", metricsRoleBindingName,
				"--clusterrole=signoz-alert-operator-metrics-reader",
				fmt.Sprintf("--serviceaccount=%s:%s", namespace, serviceAccountName),
			)
			_, err := utils.Run(cmd)
			Expect(err).NotTo(HaveOccurred(), "Failed to create ClusterRoleBinding")

			By("validating that the metrics service is available")
			cmd = exec.Command("kubectl", "get", "service", metricsServiceName, "-n", namespace)
			_, err = utils.Run(cmd)
			Expect(err).NotTo(HaveOccurred(), "Metrics service should exist")

			By("getting the service account token")
			token, err := serviceAccountToken()
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeEmpty())

			By("waiting for the metrics endpoint to be ready")
			verifyMetricsEndpointReady := func(g Gomega) {
				cmd := exec.Command("kubectl", "get", "endpoints", metricsServiceName, "-n", namespace)
				output, err := utils.Run(cmd)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(output).To(ContainSubstring("8443"), "Metrics endpoint is not ready")
			}
			Eventually(verifyMetricsEndpointReady).Should(Succeed())

			By("verifying that the controller manager is serving the metrics server")
			verifyMetricsServerStarted := func(g Gomega) {
				cmd := exec.Command("kubectl", "logs", controllerPodName, "-n", namespace)
				output, err := utils.Run(cmd)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(output).To(ContainSubstring("controller-runtime.metrics\tServing metrics server"),
					"Metrics server not yet started")
			}
			Eventually(verifyMetricsServerStarted).Should(Succeed())

			By("creating the curl-metrics pod to access the metrics endpoint")
			cmd = exec.Command("kubectl", "run", "curl-metrics", "--restart=Never",
				"--namespace", namespace,
				"--image=curlimages/curl:latest",
				"--overrides",
				fmt.Sprintf(`{
					"spec": {
						"containers": [{
							"name": "curl",
							"image": "curlimages/curl:latest",
							"command": ["/bin/sh", "-c"],
							"args": ["curl -v -k -H 'Authorization: Bearer %s' https://%s.%s.svc.cluster.local:8443/metrics"],
							"securityContext": {
								"allowPrivilegeEscalation": false,
								"capabilities": {
									"drop": ["ALL"]
								},
								"runAsNonRoot": true,
								"runAsUser": 1000,
								"seccompProfile": {
									"type": "RuntimeDefault"
								}
							}
						}],
						"serviceAccount": "%s"
					}
				}`, token, metricsServiceName, namespace, serviceAccountName))
			_, err = utils.Run(cmd)
			Expect(err).NotTo(HaveOccurred(), "Failed to create curl-metrics pod")

			By("waiting for the curl-metrics pod to complete.")
			verifyCurlUp := func(g Gomega) {
				cmd := exec.Command("kubectl", "get", "pods", "curl-metrics",
					"-o", "jsonpath={.status.phase}",
					"-n", namespace)
				output, err := utils.Run(cmd)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(output).To(Equal("Succeeded"), "curl pod in wrong status")
			}
			Eventually(verifyCurlUp, 5*time.Minute).Should(Succeed())

			By("getting the metrics by checking curl-metrics logs")
			metricsOutput := getMetricsOutput()
			Expect(metricsOutput).To(ContainSubstring(
				"controller_runtime_reconcile_total",
			))
		})

		// +kubebuilder:scaffold:e2e-webhooks-checks
	})

	// ------------------------------------------------------------------------
	// Alert lifecycle against a real SigNoz instance running in docker-compose
	// on the host. The operator reaches it via host.k3d.internal:8080 (set by
	// SIGNOZ_BASE_URL env from the Makefile).
	//
	// Auth flow (per test/e2e/casting.yaml.tmpl + utils.BootstrapSigNozAPIKey):
	//   1. SigNoz starts with a known root user/org via env vars
	//   2. BootstrapSigNozAPIKey logs in, mints a service-account API key
	//   3. That key goes into a K8s Secret, referenced from the Endpoint CR
	//
	// All cleanup is wired via DeferCleanup so it runs in LIFO order even on
	// panic, BeforeAll failure, or mid-spec abort. Finalizers are force-patched
	// before delete so a broken reconcile never causes a finalizer hang.
	// ------------------------------------------------------------------------
	Context("Alert lifecycle against real SigNoz", func() {
		const (
			alertNS       = "signoz-alert-operator-e2e"
			endpointName  = "signoz-prod"
			alertName     = "e2e-canary"
			channelName   = "e2e-null-webhook"
			secretName    = "signoz-api-key"
			secretKey     = "api-key"
			adminEmail    = "admin@e2e.local"
			adminPassword = "E2eAdmin#42Local!" //nolint:gosec // synthetic e2e credential matching test/e2e/casting.yaml.tmpl
			adminOrgID    = "01923b8a-b8a8-7e88-a06a-2c5d7a44a3e1"
		)
		var (
			apiKey string
			ruleID string
		)

		BeforeAll(func() {
			By("waiting for SigNoz to be healthy")
			Expect(utils.WaitForSigNozReady(2 * time.Minute)).To(Succeed())

			By("bootstrapping a SigNoz service-account API key (login → create SA → mint key)")
			var err error
			apiKey, err = utils.BootstrapSigNozAPIKey(adminEmail, adminPassword, adminOrgID)
			Expect(err).NotTo(HaveOccurred(), "SigNoz auth bootstrap should succeed")
			Expect(apiKey).NotTo(BeEmpty())

			By("creating a notification channel (Alerts require at least one preferredChannel)")
			// Unreachable URL — the rule never fires, so this never gets called.
			Expect(utils.CreateSigNozWebhookChannel(apiKey, channelName, "http://127.0.0.1:1/null")).To(Succeed())

			// Defensive cleanup. DeferCleanup runs LIFO regardless of how the
			// suite exits — panic, BeforeAll failure, ctrl-C. Registered first
			// so it teardowns AFTER any later DeferCleanup that adds resources.
			DeferCleanup(func() {
				if CurrentSpecReport().Failed() {
					By("[diagnostic] kubectl get alert -o yaml")
					if out, err := utils.Run(exec.Command("kubectl", "get", "alert", alertName,
						"-n", alertNS, "-o", "yaml")); err == nil {
						_, _ = fmt.Fprintf(GinkgoWriter, "Alert YAML:\n%s\n", out)
					}
					By("[diagnostic] operator pod logs")
					if out, err := utils.Run(exec.Command("kubectl", "logs",
						"-l", "control-plane=controller-manager",
						"-n", namespace, "--tail=400")); err == nil {
						_, _ = fmt.Fprintf(GinkgoWriter, "Operator logs:\n%s\n", out)
					}
				}

				// Force-clear finalizers so kubectl delete never blocks.
				_, _ = utils.Run(exec.Command("kubectl", "patch", "alert", alertName, "-n", alertNS,
					"--type=merge", "-p", `{"metadata":{"finalizers":[]}}`))
				_, _ = utils.Run(exec.Command("kubectl", "delete", "alert", alertName,
					"-n", alertNS, "--ignore-not-found", "--wait=false"))
				_, _ = utils.Run(exec.Command("kubectl", "delete", "endpoint.monitoring.hmx86.cloud",
					endpointName, "-n", alertNS, "--ignore-not-found", "--wait=false"))
				_, _ = utils.Run(exec.Command("kubectl", "delete", "secret", secretName,
					"-n", alertNS, "--ignore-not-found", "--wait=false"))
				_, _ = utils.Run(exec.Command("kubectl", "delete", "ns", alertNS,
					"--ignore-not-found", "--wait=false"))
			})

			By(fmt.Sprintf("creating namespace %s", alertNS))
			_, _ = utils.Run(exec.Command("kubectl", "create", "ns", alertNS))

			By("creating Secret holding the SigNoz API key")
			secretYAML := fmt.Sprintf(`
apiVersion: v1
kind: Secret
metadata:
  name: %s
  namespace: %s
type: Opaque
stringData:
  %s: %s
`, secretName, alertNS, secretKey, apiKey)
			secretFile := filepath.Join(os.TempDir(), "e2e-secret.yaml")
			Expect(os.WriteFile(secretFile, []byte(secretYAML), 0o600)).To(Succeed())
			_, err = utils.Run(exec.Command("kubectl", "apply", "-f", secretFile))
			Expect(err).NotTo(HaveOccurred(), "Failed to apply Secret")

			By("applying the Endpoint (with secretKeyRef → the bootstrapped key)")
			endpointYAML := fmt.Sprintf(`
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Endpoint
metadata:
  name: %s
  namespace: %s
spec:
  instanceURL: %s
  secretKeyRef:
    name: %s
    key: %s
`, endpointName, alertNS, utils.SigNozBaseURL(), secretName, secretKey)
			endpointFile := filepath.Join(os.TempDir(), "e2e-endpoint.yaml")
			Expect(os.WriteFile(endpointFile, []byte(endpointYAML), 0o644)).To(Succeed())
			_, err = utils.Run(exec.Command("kubectl", "apply", "-f", endpointFile))
			Expect(err).NotTo(HaveOccurred(), "Failed to apply Endpoint")

			By("applying the Alert")
			// Spec.Rule shape mirrors a known-working production rule:
			// fravity-analytics-layer/monitoring/aws/alerts/k8s-cdc-container-restarts.json.
			// Notes on fields SigNoz actually requires (not all are documented):
			//   - version: "v5"  — required by validation
			//   - compositeQuery.queryType: "builder"  — required
			//   - condition.selectedQueryName: matches a query.name  — required
			//   - op / matchType use NUMERIC enum codes ("1" = above / at_least_once),
			//     NOT the named enums in the OpenAPI spec ("above", "at_least_once").
			// METRIC_BASED_ALERT picked because LOGS_BASED ingestion isn't running
			// in the test compose stack. Aggregation references a metric that
			// doesn't exist (k8s.e2e.never_matches) so the rule never fires.
			alertYAML := fmt.Sprintf(`
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Alert
metadata:
  name: %s
  namespace: %s
spec:
  endpointRef:
    name: %s
  rule:
    alert: e2e-canary
    alertType: METRIC_BASED_ALERT
    ruleType: threshold_rule
    version: v5
    description: "e2e canary alert (no data expected)"
    labels:
      severity: info
    evalWindow: 5m
    frequency: 1m
    condition:
      compositeQuery:
        panelType: graph
        queryType: builder
        queries:
          - type: builder_query
            spec:
              name: A
              signal: metrics
              stepInterval: 60
              aggregations:
                - metricName: k8s.e2e.never_matches
                  timeAggregation: count
                  spaceAggregation: sum
      selectedQueryName: A
      op: "1"
      target: 0
      matchType: "1"
    preferredChannels:
      - %s
`, alertName, alertNS, endpointName, channelName)
			alertFile := filepath.Join(os.TempDir(), "e2e-alert.yaml")
			Expect(os.WriteFile(alertFile, []byte(alertYAML), 0o644)).To(Succeed())
			_, err = utils.Run(exec.Command("kubectl", "apply", "-f", alertFile))
			Expect(err).NotTo(HaveOccurred(), "Failed to apply Alert")
		})

		It("populates status.RuleID after the controller reconciles", func() {
			By("waiting for status.RuleID to be populated")
			Eventually(func(g Gomega) {
				out, err := utils.Run(exec.Command("kubectl", "get", "alert", alertName,
					"-n", alertNS, "-o", "jsonpath={.status.ruleID}"))
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(out).NotTo(BeEmpty(), "status.ruleID should be set")
				ruleID = out
			}, 2*time.Minute, 2*time.Second).Should(Succeed())
		})

		It("creates a corresponding rule in SigNoz with the right k8s_id label", func() {
			Expect(ruleID).NotTo(BeEmpty(), "previous spec must have set ruleID")

			By(fmt.Sprintf("GET %s/api/v2/rules/%s", utils.SigNozHostURL(), ruleID))
			req, _ := http.NewRequest(http.MethodGet,
				fmt.Sprintf("%s/api/v2/rules/%s", utils.SigNozHostURL(), ruleID), nil)
			req.Header.Set("SigNoz-Api-Key", apiKey)
			resp, err := http.DefaultClient.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer func() { _ = resp.Body.Close() }()
			Expect(resp.StatusCode).To(Equal(http.StatusOK), "rule should exist in SigNoz")

			var got struct {
				Data struct {
					ID     string            `json:"id"`
					Labels map[string]string `json:"labels"`
				} `json:"data"`
			}
			Expect(json.NewDecoder(resp.Body).Decode(&got)).To(Succeed())
			Expect(got.Data.ID).To(Equal(ruleID))
			Expect(got.Data.Labels).To(HaveKeyWithValue("k8s_id", alertNS+"-"+alertName),
				"controller must stamp k8s_id label for cross-cluster idempotency")
		})

		It("removes the rule from SigNoz when Alert is deleted (no finalizer hang)", func() {
			Expect(ruleID).NotTo(BeEmpty())

			// Strict 30s timeout — any longer means the operator's DELET path
			// is broken and the finalizer-removal isn't running.
			By("deleting the Alert with strict 30s --wait=true timeout (detects finalizer hangs)")
			start := time.Now()
			_, err := utils.Run(exec.Command("kubectl", "delete", "alert", alertName,
				"-n", alertNS, "--wait=true", "--timeout=30s"))
			elapsed := time.Since(start)
			Expect(err).NotTo(HaveOccurred(), "Alert delete must not hang on finalizer")
			Expect(elapsed).To(BeNumerically("<", 30*time.Second),
				"delete must complete within 30s — finalizer cleanup is broken if it hangs")

			By("verifying the rule is gone from SigNoz")
			Eventually(func(g Gomega) {
				req, _ := http.NewRequest(http.MethodGet,
					fmt.Sprintf("%s/api/v2/rules/%s", utils.SigNozHostURL(), ruleID), nil)
				req.Header.Set("SigNoz-Api-Key", apiKey)
				resp, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				_ = resp.Body.Close()
				g.Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			}, time.Minute, 2*time.Second).Should(Succeed())
		})
	})
})

// serviceAccountToken returns a token for the specified service account in the given namespace.
// It uses the Kubernetes TokenRequest API to generate a token by directly sending a request
// and parsing the resulting token from the API response.
func serviceAccountToken() (string, error) {
	const tokenRequestRawString = `{
		"apiVersion": "authentication.k8s.io/v1",
		"kind": "TokenRequest"
	}`

	// Temporary file to store the token request
	secretName := fmt.Sprintf("%s-token-request", serviceAccountName)
	tokenRequestFile := filepath.Join("/tmp", secretName)
	err := os.WriteFile(tokenRequestFile, []byte(tokenRequestRawString), os.FileMode(0o644))
	if err != nil {
		return "", err
	}

	var out string
	verifyTokenCreation := func(g Gomega) {
		// Execute kubectl command to create the token
		cmd := exec.Command("kubectl", "create", "--raw", fmt.Sprintf(
			"/api/v1/namespaces/%s/serviceaccounts/%s/token",
			namespace,
			serviceAccountName,
		), "-f", tokenRequestFile)

		output, err := cmd.CombinedOutput()
		g.Expect(err).NotTo(HaveOccurred())

		// Parse the JSON output to extract the token
		var token tokenRequest
		err = json.Unmarshal(output, &token)
		g.Expect(err).NotTo(HaveOccurred())

		out = token.Status.Token
	}
	Eventually(verifyTokenCreation).Should(Succeed())

	return out, err
}

// getMetricsOutput retrieves and returns the logs from the curl pod used to access the metrics endpoint.
func getMetricsOutput() string {
	By("getting the curl-metrics logs")
	cmd := exec.Command("kubectl", "logs", "curl-metrics", "-n", namespace)
	metricsOutput, err := utils.Run(cmd)
	Expect(err).NotTo(HaveOccurred(), "Failed to retrieve logs from curl pod")
	Expect(metricsOutput).To(ContainSubstring("< HTTP/1.1 200 OK"))
	return metricsOutput
}

// tokenRequest is a simplified representation of the Kubernetes TokenRequest API response,
// containing only the token field that we need to extract.
type tokenRequest struct {
	Status struct {
		Token string `json:"token"`
	} `json:"status"`
}
