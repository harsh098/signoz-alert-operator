# signoz-alert-operator

A Kubernetes operator for managing [SigNoz](https://signoz.io) alert rules declaratively as Kubernetes resources.

> This project is an independent community operator. It is **not** affiliated with, endorsed by, or sponsored by SigNoz Inc. "SigNoz" is a trademark of SigNoz Inc., used here for descriptive purposes only.

## Description

`signoz-alert-operator` makes SigNoz alert rules first-class Kubernetes objects so you can manage them with the same tooling you use for everything else in your cluster — GitOps, kustomize/helm overlays, RBAC, kubectl. The operator reconciles `Alert` custom resources against a SigNoz instance's `/api/v2/rules` API, keeping each rule in lockstep with its YAML manifest.

It ships two CRDs:

- **`Endpoint`** — points at a SigNoz instance (`spec.instanceURL`) and references a Kubernetes `Secret` (`spec.secretKeyRef`) holding the SigNoz API key. No controller; it's purely a reference object so individual `Alert` resources don't repeat credentials.
- **`Alert`** — references an `Endpoint` and carries the rule body as a `runtime.RawExtension` under `spec.rule`. The raw JSON is forwarded verbatim to SigNoz, so the operator stays compatible with any SigNoz alert-rule version your instance accepts.

Notable design choices:

- **Cross-cluster idempotency.** The controller stamps `labels.k8s_id = "<namespace>-<name>"` on every rule it creates in SigNoz. When the same manifest is applied to a fresh cluster pointing at the same SigNoz, the controller discovers the existing rule by that label and adopts it instead of creating a duplicate — handy for cluster migrations and management-cluster topologies.
- **Status surfaces real errors.** `status.ruleID` is the SigNoz-assigned id; `status.httpStatus` and `status.errors` carry whatever SigNoz returned, so `kubectl describe alert` is the first place to look when something's wrong.
- **Finalizer-driven cleanup.** Deleting an `Alert` CR triggers a `DELETE /api/v2/rules/{id}` to SigNoz before the object is GC'd from etcd. The operator tolerates 404 (already gone) as a no-op.

## Getting Started

### Prerequisites

- Go v1.24.0+
- Docker 17.03+
- kubectl v1.11.3+
- A Kubernetes v1.11.3+ cluster
- A reachable SigNoz instance (self-hosted or SigNoz Cloud) with a service-account API key

### Quick install from the latest release

```sh
kubectl apply -f https://github.com/harsh098/signoz-alert-operator/releases/latest/download/install.yaml
```

This applies CRDs + the operator Deployment to the `signoz-alert-operator-system` namespace using the released container image from `ghcr.io/harsh098/signoz-alert-operator`.

### Build and deploy from source

**Build and push your image:**

```sh
make docker-build docker-push IMG=ghcr.io/harsh098/signoz-alert-operator:dev
```

**Install the CRDs:**

```sh
make install
```

**Deploy the Manager:**

```sh
make deploy IMG=ghcr.io/harsh098/signoz-alert-operator:dev
```

> **NOTE**: If you encounter RBAC errors, you may need cluster-admin privileges.

**Apply the sample resources:**

```sh
kubectl apply -k config/samples/
```

The `config/samples/` directory ships an `Endpoint` (with a placeholder Secret — replace `REPLACE_ME` with your real API key) and an `Alert` targeting it.

### Uninstall

```sh
kubectl delete -k config/samples/   # delete the sample CRs
make uninstall                      # delete the CRDs
make undeploy                       # remove the controller
```

## Project Distribution

### Tag-based releases (preferred)

Push a `vX.Y.Z` tag to trigger the release workflow (see `.github/workflows/release.yml`). The workflow:

1. Builds a multi-arch (`linux/amd64`, `linux/arm64`) image and pushes to `ghcr.io/harsh098/signoz-alert-operator:vX.Y.Z` and `:latest`.
2. Generates a self-contained `dist/install.yaml` with that image tag baked in.
3. Creates a GitHub Release with `install.yaml` attached and auto-generated release notes.

Consumers install via the release-attached YAML:

```sh
kubectl apply -f https://github.com/harsh098/signoz-alert-operator/releases/download/v0.1.0/install.yaml
```

### Local install manifest

If you want the same manifest without cutting a release:

```sh
make build-installer IMG=ghcr.io/harsh098/signoz-alert-operator:dev
# → dist/install.yaml
```

### Helm chart (future)

Not yet provided. The operator-sdk Helm plugin can scaffold one:

```sh
operator-sdk edit --plugins=helm/v1-alpha
```

## Development

### Tests

- **`make test`** — unit + envtest specs against a local `kube-apiserver` + etcd (downloaded by `setup-envtest`). Fast; runs on every PR.
- **`make test-e2e`** — full e2e: creates a `k3d` cluster, brings up a real SigNoz instance via `docker compose` (pinned to v0.122.0), deploys the operator, exercises the Alert lifecycle including the delete-hangup case. Needs Docker and k3d locally; ~6 min runtime.

### Lint

- **`make lint`** — runs a custom-built `golangci-lint` binary with the [`logcheck`](https://github.com/kubernetes/utils/tree/master/logtools/logcheck) plugin enabled (see `.custom-gcl.yml`). First invocation builds the custom binary; subsequent runs reuse it.
- **`make lint-config`** — schema-validates `.golangci.yml`.

### Regenerating the SigNoz client

The SigNoz API client at `internal/signozclient/zz_generated.go` is produced by `oapi-codegen` from `hack/signoz-openapi.yaml` (vendored from SigNoz OSS at v0.122.0). To refresh:

1. Replace `hack/signoz-openapi.yaml` with the upstream copy at the desired SigNoz version (update the header block too).
2. If you need to call new SigNoz endpoints, append their `operationId`s to `hack/signoz-openapi-cfg.yaml` under `include-operation-ids`.
3. `make generate-client`.

## Contributing

Contributions are welcome. A few ground rules to keep things smooth:

1. **Open an issue first** for non-trivial changes (new CRDs, breaking schema changes, new external dependencies). For bug fixes and small improvements, a PR straight off `main` is fine.
2. **Branch and PR.** Fork the repo, create a feature branch, push, open a PR. Keep the diff focused — one logical change per PR.
3. **Before pushing**:
   - `make test` (envtest passes)
   - `make lint` (no findings)
   - If you touched controller logic, add or update a spec in `internal/controller/alert_controller_test.go`.
   - If you re-vendor the OpenAPI spec or change codegen, run `make generate-client` and commit the result.
4. **Commit style.** Concise present-tense subject; reference issues with `Fixes #N` in the body when applicable. Sign your commits if you can (`git commit -s`).
5. **Be kind in reviews.** This is a side project; expect human-paced response times.

Areas where help is especially welcome: Helm chart, more `Alert` lifecycle test scenarios, support for SigNoz dashboard resources.

**NOTE:** Run `make help` for the full list of `make` targets. Project layout follows [Kubebuilder](https://book.kubebuilder.io/introduction.html) conventions.

## License

Copyright 2026 Harsh Mishra.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Third-party attributions (including the vendored SigNoz OpenAPI spec) are listed in [`NOTICE`](NOTICE).
