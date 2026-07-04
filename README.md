# signoz-alert-operator

[![lint](https://github.com/harsh098/signoz-alert-operator/actions/workflows/lint.yml/badge.svg)](https://github.com/harsh098/signoz-alert-operator/actions/workflows/lint.yml)
[![test](https://github.com/harsh098/signoz-alert-operator/actions/workflows/test.yml/badge.svg)](https://github.com/harsh098/signoz-alert-operator/actions/workflows/test.yml)
[![e2e](https://github.com/harsh098/signoz-alert-operator/actions/workflows/test-e2e.yml/badge.svg)](https://github.com/harsh098/signoz-alert-operator/actions/workflows/test-e2e.yml)
[![release](https://img.shields.io/github/v/release/harsh098/signoz-alert-operator?include_prereleases&sort=semver)](https://github.com/harsh098/signoz-alert-operator/releases)
[![license](https://img.shields.io/github/license/harsh098/signoz-alert-operator?)](LICENSE)

A Kubernetes operator for managing [SigNoz](https://signoz.io) alert rules declaratively as Kubernetes resources.

> This project is an independent community operator. It is **not** affiliated with, endorsed by, or sponsored by SigNoz Inc. "SigNoz" is a trademark of SigNoz Inc., used here for descriptive purposes only.

---

- [What it does](#what-it-does)
- [Quick start](#quick-start)
- [Compatibility](#compatibility)
- [Versioning](#versioning)
- [License](#license)

Contributing or hacking on the operator? See **[CONTRIBUTING.md](CONTRIBUTING.md)**.

End-to-end walkthrough with a real example? See **[Usage.md](Usage.md)**.

## What it does

`signoz-alert-operator` makes SigNoz alert rules first-class Kubernetes objects. The operator reconciles `Alert` custom resources against a SigNoz instance's `/api/v2/rules` API, keeping each rule in lockstep with its YAML manifest. Manage alerts the same way you manage everything else in the cluster — GitOps, kustomize/Helm overlays, RBAC, kubectl.

Two CRDs:

- **`Endpoint`** — points at a SigNoz instance (`spec.instanceURL`) and references a Kubernetes `Secret` (`spec.secretKeyRef`) holding the SigNoz API key. No controller; just a typed reference object so `Alert` resources don't repeat credentials.
- **`Alert`** — references an `Endpoint` and carries the rule body as raw JSON under `spec.rule`. The body is forwarded verbatim to SigNoz, so the operator stays compatible with whatever alert-rule format your SigNoz version accepts. `kubectl describe alert` surfaces the SigNoz-assigned rule id, last HTTP status, and any error from the API.

## Quick start

**Install the operator** — CRDs + controller, multi-arch image, namespace `signoz-alert-operator-system`:

```sh
kubectl apply -f https://github.com/harsh098/signoz-alert-operator/releases/latest/download/install.yaml
```

**Create an Endpoint and an Alert.** Minimal shape:

```yaml
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Endpoint
metadata: { name: signoz, namespace: monitoring }
spec:
  instanceURL: https://your-tenant.us.signoz.cloud
  secretKeyRef: { name: signoz-api-key, key: api-key }
---
apiVersion: monitoring.hmx86.cloud/v1alpha1
kind: Alert
metadata: { name: my-alert, namespace: monitoring }
spec:
  endpointRef: { name: signoz }
  rule: { ... }   # raw SigNoz rule JSON
```

A complete working example (with the rule body, channel setup, troubleshooting, and the multi-cluster `k8s_id` adoption story) is in **[Usage.md](Usage.md)**.

## Compatibility

The operator's tag mirrors the SigNoz API version it was verified against. See [Versioning](#versioning) for the scheme.

**Supported SigNoz versions:** `v0.122.0` – `v0.131.1` (and their patch releases). The `/api/v2/rules` endpoints and request body (`RuletypesPostableRule`) are stable across this range — the operator's `runtime.RawExtension`-typed `spec.rule` forwards the body verbatim, so as long as your SigNoz accepts the rule, the operator can manage it. Note: SigNoz's **v5 builder-query schema tightened `signal` and query `type` into required discriminators** between v0.127 and v0.131 — rule bodies targeting v0.131 must set them (see [Usage.md](Usage.md)); this is a SigNoz-side change, the operator forwards verbatim either way.

| Operator | SigNoz API (verified against) | Kubernetes | Notes |
|---|---|---|---|
| **v1.131.0** | v0.122.0 – v0.131.1 | 1.30+ | Latest. Verified against v0.131.1 (client regen + rule round-trip + unit tests). SigNoz's v5 builder-query schema now requires `signal` and query `type` discriminators (landed between v0.127 and v0.131); rule bodies must set them, the operator forwards verbatim so no caller change. |
| **v1.127.0** | v0.122.0 – v0.127.0 | 1.30+ | e2e run against v0.127.0; wire-compat verified across the full range — only `ErrorsJSON` gained additive optional fields between v0.124 and v0.127, no caller impact. |
| **v1.124.1** | v0.124.0 | 1.30+ | Bugfix: `kubectl delete alert` no longer hangs after the SigNoz-side rule was removed via the UI. |
| **v1.124.0** | v0.124.0 | 1.30+ | First v0.124-line release. The new `/api/v2/infra_monitoring/*` endpoints in SigNoz don't affect the operator. |
| **v1.122.0** | v0.122.0 | 1.30+ | Initial release. |

The "verified against" column reflects what each tagged release was exercised against in CI/e2e — the latest operator is the recommended pick for any supported SigNoz version. Mismatched pairs outside the supported range may use endpoints one side doesn't expose; pin to a matching pair if you want predictable behaviour.

**Kubernetes:** built against controller-runtime targeting Kubernetes 1.33. Should work on any K8s 1.30+ since we only use plain CRD + Secret + ClusterRole primitives. The CEL validation on `Endpoint.spec.instanceURL` requires K8s 1.25+.

## Versioning

The operator uses a **mirror-version** scheme: an operator tag `v1.X.Y` corresponds to SigNoz upstream `v0.X.0`. The `1` prefix marks this as the operator's v1 line; `X` tracks SigNoz; `Y` is the operator's own patch counter (wire-compat doesn't move on a patch bump).

Read `v1.131.0` as: "operator v1 line, verified against SigNoz v0.131.x, initial release for that line."

This keeps the support story self-documenting — a user on SigNoz v0.131.x knows to pull operator `v1.131.x` — and lets us ship operator-only bugfixes without confusion about which SigNoz version they target.

## License

Copyright 2026 Harsh Mishra.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Third-party attributions (including the vendored SigNoz OpenAPI spec) are listed in [`NOTICE`](NOTICE).
