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

package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2" // nolint:revive,staticcheck
)

const (
	prometheusOperatorVersion = "v0.77.1"
	prometheusOperatorURL     = "https://github.com/prometheus-operator/prometheus-operator/" +
		"releases/download/%s/bundle.yaml"

	certmanagerVersion = "v1.16.3"
	certmanagerURLTmpl = "https://github.com/cert-manager/cert-manager/releases/download/%s/cert-manager.yaml"
)

func warnError(err error) {
	_, _ = fmt.Fprintf(GinkgoWriter, "warning: %v\n", err)
}

// Run executes the provided command within this context
func Run(cmd *exec.Cmd) (string, error) {
	dir, _ := GetProjectDir()
	cmd.Dir = dir

	if err := os.Chdir(cmd.Dir); err != nil {
		_, _ = fmt.Fprintf(GinkgoWriter, "chdir dir: %q\n", err)
	}

	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	command := strings.Join(cmd.Args, " ")
	_, _ = fmt.Fprintf(GinkgoWriter, "running: %q\n", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("%q failed with error %q: %w", command, string(output), err)
	}

	return string(output), nil
}

// InstallPrometheusOperator installs the prometheus Operator to be used to export the enabled metrics.
func InstallPrometheusOperator() error {
	url := fmt.Sprintf(prometheusOperatorURL, prometheusOperatorVersion)
	cmd := exec.Command("kubectl", "create", "-f", url)
	_, err := Run(cmd)
	return err
}

// UninstallPrometheusOperator uninstalls the prometheus
func UninstallPrometheusOperator() {
	url := fmt.Sprintf(prometheusOperatorURL, prometheusOperatorVersion)
	cmd := exec.Command("kubectl", "delete", "-f", url)
	if _, err := Run(cmd); err != nil {
		warnError(err)
	}
}

// IsPrometheusCRDsInstalled checks if any Prometheus CRDs are installed
// by verifying the existence of key CRDs related to Prometheus.
func IsPrometheusCRDsInstalled() bool {
	// List of common Prometheus CRDs
	prometheusCRDs := []string{
		"prometheuses.monitoring.coreos.com",
		"prometheusrules.monitoring.coreos.com",
		"prometheusagents.monitoring.coreos.com",
	}

	cmd := exec.Command("kubectl", "get", "crds", "-o", "custom-columns=NAME:.metadata.name")
	output, err := Run(cmd)
	if err != nil {
		return false
	}
	crdList := GetNonEmptyLines(output)
	for _, crd := range prometheusCRDs {
		for _, line := range crdList {
			if strings.Contains(line, crd) {
				return true
			}
		}
	}

	return false
}

// UninstallCertManager uninstalls the cert manager
func UninstallCertManager() {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "delete", "-f", url)
	if _, err := Run(cmd); err != nil {
		warnError(err)
	}
}

// InstallCertManager installs the cert manager bundle.
func InstallCertManager() error {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "apply", "-f", url)
	if _, err := Run(cmd); err != nil {
		return err
	}
	// Wait for cert-manager-webhook to be ready, which can take time if cert-manager
	// was re-installed after uninstalling on a cluster.
	cmd = exec.Command("kubectl", "wait", "deployment.apps/cert-manager-webhook",
		"--for", "condition=Available",
		"--namespace", "cert-manager",
		"--timeout", "5m",
	)

	_, err := Run(cmd)
	return err
}

// IsCertManagerCRDsInstalled checks if any Cert Manager CRDs are installed
// by verifying the existence of key CRDs related to Cert Manager.
func IsCertManagerCRDsInstalled() bool {
	// List of common Cert Manager CRDs
	certManagerCRDs := []string{
		"certificates.cert-manager.io",
		"issuers.cert-manager.io",
		"clusterissuers.cert-manager.io",
		"certificaterequests.cert-manager.io",
		"orders.acme.cert-manager.io",
		"challenges.acme.cert-manager.io",
	}

	// Execute the kubectl command to get all CRDs
	cmd := exec.Command("kubectl", "get", "crds")
	output, err := Run(cmd)
	if err != nil {
		return false
	}

	// Check if any of the Cert Manager CRDs are present
	crdList := GetNonEmptyLines(output)
	for _, crd := range certManagerCRDs {
		for _, line := range crdList {
			if strings.Contains(line, crd) {
				return true
			}
		}
	}

	return false
}

// LoadImageToK3dClusterWithName loads a local docker image into the k3d cluster.
// Cluster name comes from $K3D_CLUSTER, defaulting to "k3s-default".
func LoadImageToK3dClusterWithName(name string) error {
	cluster := "k3s-default"
	if v, ok := os.LookupEnv("K3D_CLUSTER"); ok {
		cluster = v
	}
	cmd := exec.Command("k3d", "image", "import", name, "-c", cluster)
	_, err := Run(cmd)
	return err
}

// SigNozBaseURL returns the URL the operator (inside k3d) should use to reach
// SigNoz. Defaults to host.k3d.internal:8080 so pods can reach the host's
// docker-compose SigNoz. Override with $SIGNOZ_BASE_URL.
func SigNozBaseURL() string {
	if v, ok := os.LookupEnv("SIGNOZ_BASE_URL"); ok {
		return v
	}
	return "http://host.k3d.internal:8080"
}

// SigNozHostURL returns the URL usable from the host (the test process) to
// reach SigNoz directly — used for readiness checks and post-test verification.
// Override with $SIGNOZ_HOST_URL.
func SigNozHostURL() string {
	if v, ok := os.LookupEnv("SIGNOZ_HOST_URL"); ok {
		return v
	}
	return "http://localhost:8080"
}

// WaitForSigNozReady polls SigNoz /api/v1/health from the host until 200 or
// the timeout elapses. /api/v1/health is the only anonymous-OK endpoint; the
// rules API requires a service-account key (see BootstrapSigNozAPIKey).
func WaitForSigNozReady(timeout time.Duration) error {
	url := SigNozHostURL() + "/api/v1/health"
	deadline := time.Now().Add(timeout)
	client := &http.Client{Timeout: 2 * time.Second}
	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		if err == nil {
			_ = resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("signoz not ready at %s within %s", url, timeout)
}

// BootstrapSigNozAPIKey mints a service-account API key on a fresh SigNoz
// instance via the public API. Requires that the root user has been
// provisioned via SIGNOZ_USER_ROOT_* env vars (see test/e2e/casting.yaml.tmpl).
//
// Five HTTP calls:
//  1. POST /api/v2/sessions/email_password  → JWT (uses email/password/orgID)
//  2. GET  /api/v1/roles                    → find ADMIN role id (Bearer JWT)
//  3. POST /api/v1/service_accounts         → create SA, get its id
//  4. POST /api/v1/service_accounts/{id}/roles → attach ADMIN role
//  5. POST /api/v1/service_accounts/{id}/keys  → mint API key
//
// The returned string is the API key value to put in the SigNoz-Api-Key header.
func BootstrapSigNozAPIKey(email, password, orgID string) (string, error) {
	base := SigNozHostURL()
	hc := &http.Client{Timeout: 15 * time.Second}

	jwt, err := signozLogin(hc, base, email, password, orgID)
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}

	adminRoleID, err := signozFindAdminRoleID(hc, base, jwt)
	if err != nil {
		return "", fmt.Errorf("find admin role: %w", err)
	}

	saID, err := signozCreateServiceAccount(hc, base, jwt, "e2e-operator")
	if err != nil {
		return "", fmt.Errorf("create service account: %w", err)
	}

	if err := signozAssignRoleToSA(hc, base, jwt, saID, adminRoleID); err != nil {
		return "", fmt.Errorf("assign admin role: %w", err)
	}

	key, err := signozCreateAPIKey(hc, base, jwt, saID, "e2e-key")
	if err != nil {
		return "", fmt.Errorf("create api key: %w", err)
	}
	return key, nil
}

// CreateSigNozWebhookChannel creates a webhook notification channel via the
// SigNoz alertmanager API. Alerts must reference at least one channel by name
// in their `preferredChannels`, so e2e tests provision a dummy channel here
// pointing at an unreachable URL (the rule never fires anyway).
//
// Auth uses the API key minted by BootstrapSigNozAPIKey (sent via
// SigNoz-Api-Key header — the channel API accepts both that and Bearer JWTs).
func CreateSigNozWebhookChannel(apiKey, name, webhookURL string) error {
	body := map[string]any{
		"name": name,
		"webhook_configs": []map[string]any{
			{
				"url":           webhookURL,
				"send_resolved": false,
			},
		},
	}
	buf, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, SigNozHostURL()+"/api/v1/channels", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("SigNoz-Api-Key", apiKey)
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return fmt.Errorf("create channel: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	// 201 on success; 409 if a channel with this name already exists from a
	// prior failed run that wasn't cleaned up — treat as success.
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusConflict {
		return nil
	}
	return fmt.Errorf("create channel: status=%d body=%s", resp.StatusCode, string(respBody))
}

// signozEnvelope unwraps the {"status":"...","data":...} response wrapper.
type signozEnvelope struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func signozPostJSON(hc *http.Client, url, jwt string, body any) (*signozEnvelope, error) {
	buf, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("status=%d body=%s", resp.StatusCode, string(respBody))
	}
	// 204 No Content (and rare 200 with empty body) — return a zero envelope
	// instead of trying to decode "". Callers that need .Data must check.
	if len(bytes.TrimSpace(respBody)) == 0 {
		return &signozEnvelope{}, nil
	}
	var env signozEnvelope
	if err := json.Unmarshal(respBody, &env); err != nil {
		return nil, fmt.Errorf("decode envelope: %w (body=%s)", err, string(respBody))
	}
	return &env, nil
}

func signozGetJSON(hc *http.Client, url, jwt string) (*signozEnvelope, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("status=%d body=%s", resp.StatusCode, string(respBody))
	}
	var env signozEnvelope
	if err := json.Unmarshal(respBody, &env); err != nil {
		return nil, fmt.Errorf("decode envelope: %w (body=%s)", err, string(respBody))
	}
	return &env, nil
}

func signozLogin(hc *http.Client, base, email, password, orgID string) (string, error) {
	env, err := signozPostJSON(hc, base+"/api/v2/sessions/email_password", "",
		map[string]string{"email": email, "password": password, "orgId": orgID})
	if err != nil {
		return "", err
	}
	var tok struct {
		AccessToken string `json:"accessToken"`
	}
	if err := json.Unmarshal(env.Data, &tok); err != nil {
		return "", fmt.Errorf("decode token: %w", err)
	}
	if tok.AccessToken == "" {
		return "", fmt.Errorf("login response missing accessToken: %s", string(env.Data))
	}
	return tok.AccessToken, nil
}

func signozFindAdminRoleID(hc *http.Client, base, jwt string) (string, error) {
	env, err := signozGetJSON(hc, base+"/api/v1/roles", jwt)
	if err != nil {
		return "", err
	}
	var roles []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(env.Data, &roles); err != nil {
		return "", fmt.Errorf("decode roles: %w", err)
	}
	// Managed role names in SigNoz are "signoz-admin", "signoz-editor",
	// "signoz-viewer", "signoz-anonymous" — see pkg/types/coretypes/registry_managed_role.go.
	for _, r := range roles {
		if strings.EqualFold(r.Name, "signoz-admin") {
			return r.ID, nil
		}
	}
	names := make([]string, len(roles))
	for i, r := range roles {
		names[i] = r.Name
	}
	return "", fmt.Errorf("no signoz-admin role found (got: %s)", strings.Join(names, ", "))
}

func signozCreateServiceAccount(hc *http.Client, base, jwt, name string) (string, error) {
	env, err := signozPostJSON(hc, base+"/api/v1/service_accounts", jwt,
		map[string]string{"name": name})
	if err != nil {
		return "", err
	}
	var ident struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(env.Data, &ident); err != nil {
		return "", fmt.Errorf("decode service account: %w", err)
	}
	if ident.ID == "" {
		return "", fmt.Errorf("service account response missing id: %s", string(env.Data))
	}
	return ident.ID, nil
}

func signozAssignRoleToSA(hc *http.Client, base, jwt, saID, roleID string) error {
	_, err := signozPostJSON(hc,
		fmt.Sprintf("%s/api/v1/service_accounts/%s/roles", base, saID), jwt,
		map[string]string{"id": roleID})
	return err
}

func signozCreateAPIKey(hc *http.Client, base, jwt, saID, name string) (string, error) {
	// expiresAt is uint64 — seconds since epoch (1 year from now is plenty for tests).
	expiresAt := uint64(time.Now().Add(365 * 24 * time.Hour).Unix())
	env, err := signozPostJSON(hc,
		fmt.Sprintf("%s/api/v1/service_accounts/%s/keys", base, saID), jwt,
		map[string]any{"name": name, "expiresAt": expiresAt})
	if err != nil {
		return "", err
	}
	var out struct {
		Key string `json:"key"`
	}
	if err := json.Unmarshal(env.Data, &out); err != nil {
		return "", fmt.Errorf("decode api key: %w", err)
	}
	if out.Key == "" {
		return "", fmt.Errorf("api key response missing key: %s", string(env.Data))
	}
	return out.Key, nil
}

// GetNonEmptyLines converts given command output string into individual objects
// according to line breakers, and ignores the empty elements in it.
func GetNonEmptyLines(output string) []string {
	var res []string
	for element := range strings.SplitSeq(output, "\n") {
		if element != "" {
			res = append(res, element)
		}
	}

	return res
}

// GetProjectDir will return the directory where the project is
func GetProjectDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, fmt.Errorf("failed to get current working directory: %w", err)
	}
	wd = strings.ReplaceAll(wd, "/test/e2e", "")
	return wd, nil
}

// UncommentCode searches for target in the file and remove the comment prefix
// of the target content. The target content may span multiple lines.
func UncommentCode(filename, target, prefix string) error {
	// false positive
	// nolint:gosec
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", filename, err)
	}
	strContent := string(content)

	idx := strings.Index(strContent, target)
	if idx < 0 {
		return fmt.Errorf("unable to find the code %q to be uncomment", target)
	}

	out := new(bytes.Buffer)
	_, err = out.Write(content[:idx])
	if err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	scanner := bufio.NewScanner(bytes.NewBufferString(target))
	if !scanner.Scan() {
		return nil
	}
	for {
		if _, err = out.WriteString(strings.TrimPrefix(scanner.Text(), prefix)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
		// Avoid writing a newline in case the previous line was the last in target.
		if !scanner.Scan() {
			break
		}
		if _, err = out.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	if _, err = out.Write(content[idx+len(target):]); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	// false positive
	// nolint:gosec
	if err = os.WriteFile(filename, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %q: %w", filename, err)
	}

	return nil
}
