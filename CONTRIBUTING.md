# Contributing

Thanks for hacking on `signoz-alert-operator`. This file is for people working on the operator itself — design, dev loop, tests, release process. End-user docs live in [README.md](README.md) and [Usage.md](Usage.md).

- [Design](#design)
- [Development loop](#development-loop)
- [Tests](#tests)
- [Lint](#lint)
- [Regenerating the SigNoz client](#regenerating-the-signoz-client)
- [Releases](#releases)
- [Versioning policy](#versioning-policy)
- [PR ground rules](#pr-ground-rules)
- [Open work](#open-work)

## Design

Decisions worth knowing about before editing the controller:

- **Cross-cluster idempotency via `k8s_id` label.** Every rule the operator creates in SigNoz gets `labels.k8s_id = "<namespace>-<name>"`. When the same manifest is later applied to a fresh cluster pointing at the same SigNoz, the controller lists rules, finds one tagged with the matching `k8s_id`, and **adopts** it — `status.ruleID` gets populated from the existing rule and the operator PUTs an update instead of creating a duplicate. Built for cluster migrations and management-cluster topologies.
- **`spec.rule` is a `runtime.RawExtension`** with `+kubebuilder:validation:Schemaless` + `+kubebuilder:pruning:PreserveUnknownFields`. The body is forwarded verbatim to SigNoz, so the CRD doesn't need to track SigNoz's internal schema churn. Validation errors surface via `status.errors`.
- **Status is the error surface.** `status.ruleID` is the SigNoz-assigned id; `status.httpStatus` and `status.errors` carry the most recent API call's outcome. `kubectl describe alert` is the first place to look when something's wrong.
- **Finalizer-driven cleanup.** Deleting an `Alert` CR triggers `DELETE /api/v2/rules/{id}` before the K8s object is GC'd. Both 204 (deleted) and 404 (already gone via UI) count as cleanup-success; the finalizer comes off and the CR is reaped. Other status codes hold the finalizer for retry until SigNoz cooperates.
- **No watcher on SigNoz.** Drift (someone deletes a rule via the UI) is detected only on the next K8s-side reconcile — periodic cache resync (~10h default), operator restart, or a manifest edit. Adding a `Result{RequeueAfter: 5 * time.Minute}` on the happy path would tighten this without much load; see [Open work](#open-work).

The controller is organised as a four-layer state machine: `plan()` is pure (decides what op to run from current state), `probe()` reads SigNoz state, `apply()` performs the side effect, and `Reconcile()` glues them with the K8s lifecycle (finalizer add/remove, status writes). Each layer has its own tests.

## Development loop

```sh
make help                  # full list of make targets
make test                  # unit + envtest (~10s)
make build                 # build manager binary into bin/
make run                   # run controller against ~/.kube/config
```

Project layout follows [Kubebuilder](https://book.kubebuilder.io/introduction.html) conventions:

| Path | What's there |
|---|---|
| `api/v1alpha1/` | CRD type definitions (`Alert`, `Endpoint`) and kubebuilder markers |
| `internal/controller/` | `Reconcile`, state machine (`plan`), helpers (`probe`, `apply`, `resolveClient`) |
| `internal/signozclient/` | Generated SigNoz API client (don't edit `zz_generated.go`) |
| `hack/signoz-openapi.yaml` | Vendored SigNoz OpenAPI spec |
| `hack/signoz-openapi-cfg.yaml` | oapi-codegen scope (which operations to include) |
| `test/e2e/` | k3d-based end-to-end suite |
| `config/` | kustomize bases for CRDs, RBAC, manager Deployment |

## Tests

| Target | What it does | Time |
|---|---|---|
| `make test` | Unit + envtest (`internal/controller/alert_controller_test.go`, `alert_unit_test.go`). Uses a local `kube-apiserver` + etcd via `setup-envtest`, mocks SigNoz with `httptest`. Runs on every PR. | ~10s |
| `make test-e2e` | Real end-to-end: spins up a k3d cluster + real SigNoz via `docker compose` (pinned to `SIGNOZ_VERSION`), deploys the operator from the just-built image, exercises the full Alert lifecycle. Needs Docker + k3d locally. | ~6 min |

The envtest layer is the primary safety net. The e2e layer catches integration-level issues you can't fake (auth bootstrap against a real SigNoz, the SigNoz-API-Key vs `Authorization` header dance, real CRD admission, real finalizer behaviour on a real apiserver).

## Lint

```sh
make lint            # runs custom-built golangci-lint with the logcheck plugin
make lint-config     # schema-validate .golangci.yml
```

The lint binary is built once from `.custom-gcl.yml` to pick up the [`logcheck`](https://github.com/kubernetes/utils/tree/master/logtools/logcheck) plugin (Kubernetes structured-logging conventions). Subsequent runs reuse `bin/custom-gcl`. Bumping the lint version means bumping both `GOLANGCI_LINT_VERSION` in the Makefile and the `version:` field in `.custom-gcl.yml`.

## Regenerating the SigNoz client

The SigNoz API client at `internal/signozclient/zz_generated.go` is produced by `oapi-codegen` from `hack/signoz-openapi.yaml` (vendored from SigNoz OSS). To bump to a new SigNoz version:

1. **Re-vendor the spec.** Replace `hack/signoz-openapi.yaml` with the upstream copy at the desired tag — update the header block (`Source`, `Version`, `Retrieved`) too.
2. **Diff against the previous version.** Specifically check whether endpoints we use changed shape: `/api/v2/rules`, `/api/v2/sessions/email_password`, `/api/v1/roles`, `/api/v1/service_accounts`, `/api/v1/service_accounts/{id}/keys`, `/api/v1/service_accounts/{id}/roles`, `/api/v1/channels`. Cosmetic YAML reformatting (quote style, indent shifts in `required:` lists) is safe to ignore; structural changes to request/response schemas are not.
3. **Bump `SIGNOZ_VERSION`** in the Makefile.
4. **Regenerate the client:** `make generate-client`.
5. **Run tests:** `make test` to confirm the new client still compiles and integrates.
6. **If new SigNoz endpoints need calling**, append their `operationId`s to `hack/signoz-openapi-cfg.yaml` under `include-operation-ids` *before* step 4. The generator scopes by operation id to keep the client small (~2 KLoC instead of ~32 KLoC).
7. **Bump documentation:** README compatibility matrix, Usage.md references, the comment in `api/v1alpha1/alert_types.go`.

## Releases

Push a `vX.Y.Z` tag to trigger the release workflow ([`.github/workflows/release.yml`](.github/workflows/release.yml)). The workflow:

1. Builds a multi-arch (`linux/amd64`, `linux/arm64`) image and pushes to `ghcr.io/harsh098/signoz-alert-operator:vX.Y.Z` and `:latest`.
2. Generates `dist/install.yaml` with that image tag baked in (`make build-installer`).
3. Creates a GitHub Release with `install.yaml` attached and auto-generated release notes.

```sh
# After committing the version bump:
git tag v1.X.Y && git push --tags

# For local-only manifest without cutting a release:
make build-installer IMG=ghcr.io/harsh098/signoz-alert-operator:dev
# → dist/install.yaml
```

## Versioning policy

The operator uses a mirror-version scheme: tag `v1.X.Y` corresponds to SigNoz upstream `v0.X.0`. The `1` prefix marks the operator's v1 line; `X` tracks SigNoz; `Y` is the operator's patch counter — wire-compat doesn't move on patch bumps.

When SigNoz cuts a new minor (`v0.X+1.0`):

1. Re-vendor + diff per [Regenerating the SigNoz client](#regenerating-the-signoz-client).
2. If endpoints we use are unchanged: bump `SIGNOZ_VERSION`, run `make test test-e2e`, update README compat matrix, tag `v1.X+1.0`.
3. If endpoints changed shape: the patch is bigger — adjust client callers, possibly add a new CRD field, write new tests.

Operator-only bugfixes (no wire change) bump `Y`: `v1.124.0` → `v1.124.1`.

## PR ground rules

1. **Open an issue first** for non-trivial changes (new CRDs, breaking schema changes, external dependencies). Bug fixes and small improvements: PR straight off `main` is fine.
2. **Branch and PR.** Keep diffs focused — one logical change per PR.
3. **Before pushing**:
   - `make test` passes
   - `make lint` clean
   - If you touched controller logic, add or update a spec in `internal/controller/alert_controller_test.go`.
   - If you re-vendored the OpenAPI spec or changed codegen, run `make generate-client` and commit the result.
4. **Commit style.** Concise present-tense subject; reference issues with `Fixes #N` in the body where applicable.
5. **Be kind in reviews.** This is a side project; expect human-paced response times.

## Open work

Areas where help is especially welcome:

- **Helm chart.** The operator-sdk Helm plugin scaffolds one with `operator-sdk edit --plugins=helm/v1-alpha`.
- **Per-Endpoint rate limiter.** A mass `kubectl apply` of many Alerts currently bursts SigNoz with parallel requests; a token-bucket keyed by Endpoint URL would smooth that out.
- **Drift detection.** Periodic `RequeueAfter` on happy-path Reconcile so UI-side deletions/edits are caught in minutes rather than the controller-runtime default ~10h cache resync.
- **`observedGeneration` on Status.** Avoid no-op PUTs to SigNoz when the K8s spec hasn't changed since the last successful reconcile.
- **Dashboard CRD.** The `/api/v1/dashboards` endpoints are already in the vendored spec; same controller pattern as `Alert` (POST/PUT/DELETE, finalizer cleanup, `k8s_id` label).
- **More e2e scenarios.** SigNoz 4xx/5xx error paths, multi-Endpoint reconciliation, adoption after a partial cluster migration.
