# Usage

A walkthrough of provisioning a real alert via `signoz-alert-operator`, end-to-end. Pick a SigNoz instance you can reach from the cluster, follow the four steps, and you'll have a `kubectl`-managed alert in SigNoz within a few minutes.

The running example is a **container-restart alert**: fire if any pod of a given Deployment restarts more than three times in a 5-minute window. Replace metric names and filters with whatever fits your workload.

---

## Prerequisites

1. **The operator is installed** on your cluster. From the latest release:

   ```sh
   kubectl apply -f https://github.com/harsh098/signoz-alert-operator/releases/latest/download/install.yaml
   ```

   Verify the controller pod is running:

   ```sh
   kubectl -n signoz-alert-operator-system get pods
   # signoz-alert-operator-controller-manager-...   Running
   ```

2. **A reachable SigNoz instance.** Self-hosted (e.g. on the same cluster, at `http://signoz-frontend.signoz-system.svc.cluster.local:8080`) or SigNoz Cloud (`https://<tenant>.signoz.cloud`). The operator's pods need network access to whatever URL you use.

3. **A SigNoz API key** scoped at least to `signoz-admin` (alert rule mutation is admin-gated).

   - **SigNoz Cloud:** Settings → Workspace Settings → API Keys → New key.
   - **Self-hosted:** Log in as an admin, create a service account in Settings → API Keys, then create a key for it. You can also script this end-to-end against `/api/v2/sessions/email_password` → `/api/v1/service_accounts` → `/api/v1/service_accounts/{id}/keys`; see `test/utils/utils.go:BootstrapSigNozAPIKey` for a working reference.

4. **At least one notification channel configured in SigNoz.** SigNoz refuses to create a rule without `preferredChannels` referencing an existing channel by name. Create one in the SigNoz UI (Settings → Channels) — Slack, email, webhook, PagerDuty, whatever — and remember the **name**.

---

## Step 1 — Store the API key in a Secret

The Secret lives in the **same namespace** where you'll create the `Endpoint`. Adjust the namespace as needed:

```yaml
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: signoz-prod-credentials
  namespace: monitoring
type: Opaque
stringData:
  api-key: <YOUR_SIGNOZ_API_KEY>
```

```sh
kubectl apply -f secret.yaml
```

> Want this managed by external-secrets / sealed-secrets / sops? Go ahead — the operator only cares that a `Secret` with the named key exists.

---

## Step 2 — Declare the SigNoz target via an `Endpoint`

```yaml
# endpoint.yaml
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Endpoint
metadata:
  name: signoz-prod
  namespace: monitoring
spec:
  # SigNoz Cloud:
  instanceURL: https://your-tenant.signoz.cloud
  # Or self-hosted:
  # instanceURL: http://signoz-frontend.signoz-system.svc.cluster.local:8080
  secretKeyRef:
    name: signoz-prod-credentials
    key: api-key
```

```sh
kubectl apply -f endpoint.yaml
```

The `Endpoint` has no controller — it's purely a typed reference object so individual `Alert` resources don't have to repeat the URL + Secret pointer. Many `Alert`s in the same namespace can target the same `Endpoint`.

---

## Step 3 — Define the `Alert`

```yaml
# alert.yaml
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Alert
metadata:
  name: cdc-server-container-restarts
  namespace: monitoring
spec:
  endpointRef:
    name: signoz-prod        # Endpoint in the same namespace

  # Everything under spec.rule is forwarded verbatim to SigNoz POST /api/v2/rules.
  # Shape mirrors what SigNoz's UI produces — easiest way to bootstrap a new rule
  # is to build it in the UI, export the JSON, and paste it here.
  rule:
    alert: "CDC Server container restarts"
    alertType: METRIC_BASED_ALERT
    ruleType: threshold_rule
    version: v5
    description: |
      CDC Server pods are restarting too frequently. This usually indicates
      OOM kills, liveness probe failures, or panics in the workload.
    evalWindow: 5m
    frequency: 1m
    labels:
      severity: warning
      team: data-platform
      service: cdc-server
    annotations:
      summary: "{{$labels.k8s_pod_name}} has restarted {{$value}} times in 5m"
      runbook: "https://runbooks.example.com/cdc-server-restarts"
    condition:
      compositeQuery:
        queryType: builder
        panelType: graph
        queries:
          - type: builder_query
            spec:
              name: A
              signal: metrics
              stepInterval: 60
              aggregations:
                - metricName: k8s.container.restarts
                  timeAggregation: increase
                  spaceAggregation: sum
              filter:
                expression: "k8s.container.name = 'cdc-server'"
              groupBy:
                - name: k8s.pod.name
                  fieldContext: resource
      selectedQueryName: A
      op: "1"          # 1 = above; SigNoz uses numeric codes here, NOT the spec's named enums
      target: 3
      matchType: "1"   # 1 = at_least_once
    preferredChannels:
      - slack-data-platform   # must match a channel name you created in SigNoz
```

```sh
kubectl apply -f alert.yaml
```

A few things worth flagging about the rule body:

- **`spec.rule` is `runtime.RawExtension`** — the CRD doesn't validate field-by-field. Anything SigNoz rejects will surface on `status.errors` rather than being caught by the API server. Pro: the operator stays compatible across SigNoz versions without code changes. Con: typos like `panelTypee` slip past `kubectl apply` and only fail at reconcile time.
- **`condition.op` and `condition.matchType` use numeric string codes** (`"1"` = above / at_least_once), not the named enums in SigNoz's public OpenAPI spec. The UI's exported JSON uses numerics, so we do too.
- **`version: v5`** is required by SigNoz's validation even though it looks optional in the spec.
- **`labels.k8s_id` will be added automatically by the controller** — don't set it yourself. It's how the operator finds the rule in SigNoz across cluster moves.

---

## Step 4 — Verify the alert was created

```sh
kubectl -n monitoring get alert cdc-server-container-restarts -o yaml
```

Look at `status`:

```yaml
status:
  ruleID: "1234"        # SigNoz-assigned id; empty if the rule never got created
  httpStatus: 201       # 201 on initial create, 200 on subsequent updates
  errors: ""            # SigNoz's error body if anything went wrong
```

A populated `ruleID` and `httpStatus: 201` / `200` means the rule exists in SigNoz. You can confirm in the SigNoz UI (Alerts → All Alerts → look for "CDC Server container restarts").

If `ruleID` is empty and `errors` is set, that's SigNoz rejecting your rule body. The error text usually says which field is wrong; fix `spec.rule` and re-apply.

---

## Updating a rule

Edit the `Alert` manifest and reapply:

```sh
kubectl apply -f alert.yaml
```

The controller detects the change, sends `PUT /api/v2/rules/{ruleID}` to SigNoz, and bumps `status.httpStatus` to 200. SigNoz keeps the same rule id, so existing alert history and notifications continue.

---

## Deleting a rule

```sh
kubectl -n monitoring delete alert cdc-server-container-restarts
```

The operator's finalizer ensures the SigNoz-side rule is deleted (`DELETE /api/v2/rules/{ruleID}`) *before* the K8s object is garbage-collected. A successful delete typically completes in <2 seconds; if `kubectl delete` hangs longer than that, check `kubectl describe alert ...` for the latest `status.errors` — the finalizer holds the object alive until cleanup succeeds.

If the rule was deleted from the SigNoz UI before you ran `kubectl delete`, the controller tolerates the 404 and removes the finalizer anyway.

---

## Multi-endpoint setup

A single operator instance can manage alerts across many SigNoz targets — create one `Endpoint` per target:

```yaml
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Endpoint
metadata: { name: signoz-prod, namespace: monitoring }
spec:
  instanceURL: https://prod.signoz.cloud
  secretKeyRef: { name: signoz-prod-credentials, key: api-key }
---
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Endpoint
metadata: { name: signoz-staging, namespace: monitoring }
spec:
  instanceURL: https://staging.signoz.cloud
  secretKeyRef: { name: signoz-staging-credentials, key: api-key }
```

Each `Alert` picks its target via `spec.endpointRef.name`. Convenient pattern for a management cluster that ships alerts to per-environment SigNoz instances.

---

## Cluster migration (the `k8s_id` trick)

If you `kubectl apply -k` the same alert manifests into a new cluster pointing at the same SigNoz instance:

1. The new cluster's `Alert` CR starts with empty `status.ruleID`.
2. On first reconcile, the controller lists rules in SigNoz and looks for one labelled `k8s_id: <namespace>-<name>`.
3. Match found → it adopts that rule's id into `status.ruleID` and `PUT`s an update with your latest spec. **No duplicate rule** is created in SigNoz.

This makes it safe to apply the same manifests in dev / staging / prod clusters when they share a SigNoz backend.

---

## Troubleshooting

**`status.errors` says `"endpoint <name> has no secretKeyRef set"`**
The `Endpoint` is missing `spec.secretKeyRef`. Self-hosted SigNoz still requires auth for `/api/v2/rules` — set the Secret reference.

**`status.errors` says `"secret <name> has no key <key>"`**
The Secret you referenced doesn't have the field you named. `kubectl get secret -o yaml` and check `data` (or `stringData` if you used it on apply).

**`status.errors` contains `"unauthenticated"`**
The API key is wrong, expired, or doesn't have admin scope on the SigNoz side. Re-create a key with `signoz-admin` role.

**`status.errors` contains `"validation failed"`**
SigNoz rejected the rule body. Most common causes (verified against SigNoz v0.124.0):
- Missing `version: v5`.
- Missing `compositeQuery.queryType: builder`.
- Missing `condition.selectedQueryName` (must match a `queries[].spec.name`).
- Empty `preferredChannels`, or a channel name that doesn't exist in SigNoz.
- `op` / `matchType` using named enums (`"above"`, `"at_least_once"`) instead of numeric codes (`"1"`).

**`kubectl delete alert ...` hangs forever**
The finalizer is stuck because cleanup keeps failing. `kubectl describe alert ...` → look at `status.errors`. Usually the SigNoz instance is unreachable. Fix connectivity (or, in an emergency, remove the finalizer manually: `kubectl patch alert ... --type=merge -p '{"metadata":{"finalizers":[]}}'`).

**Need to see what the operator is doing**

```sh
kubectl -n signoz-alert-operator-system logs -l control-plane=controller-manager -f
```

---

## Reference

- Sample manifests live under [`config/samples/`](config/samples/).
- The operator's design (CRDs, idempotency, status semantics) is documented in [README.md](README.md).
- Full SigNoz API reference: [signoz.io/docs](https://signoz.io/docs/userguide/).
