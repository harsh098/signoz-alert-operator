# Changelog

## [1.131.2] - 2026-07-04

### Fixed
- Docker builder base image bumped `golang:1.24` → `golang:1.26` to satisfy the
  `go 1.26.4` module directive; `make docker-build` (and therefore the e2e
  `BeforeSuite` and the release image build) no longer fails with
  `go.mod requires go >= 1.26.4 (running go 1.24.x; GOTOOLCHAIN=local)`.

Operator-only fix; no wire change and no CRD change.

## [1.131.1] - 2026-07-04

### Added
- Periodic drift reconcile: a successful reconcile now re-queues the Alert on a
  ~5-minute interval (with ±10% jitter), so drift in SigNoz (e.g. a rule edited
  or deleted in the UI) is re-asserted without needing a change to the CR.
  Transient failures keep their faster 30s retry.
- OperatorHub.io / OLM packaging: a `ClusterServiceVersion` base
  (`config/manifests/bases/`) with full OperatorHub metadata (display name,
  description, keywords, maintainer, provider, links, `Monitoring` category,
  install modes, and a placeholder icon), CSV spec/status descriptors on the
  `Alert` and `Endpoint` CRDs, and a generated OLM bundle (`bundle/`,
  `bundle.Dockerfile`). Validated against the `operatorhubv2`, `capabilities`,
  `categories`, and `good-practices` validators.

### Changed
- Bumped SigNoz API compatibility from **v0.127.0 to v0.131.1**: re-vendored the
  OpenAPI spec (`hack/signoz-openapi.yaml`) and regenerated the narrow client
  (`internal/signozclient/zz_generated.go`). The `/api/v2/rules` operations and
  the `RuletypesPostableRule` request body are unchanged; rule bodies are still
  forwarded verbatim.
- E2E: migrated the SigNoz bring-up from the vendored `deploy/` docker-compose
  tree (removed upstream in SigNoz v0.130.0) to **Foundry** (`foundryctl`
  v0.2.11). SigNoz is now rendered from a declarative `test/e2e/casting.yaml.tmpl`
  via `foundryctl forge` and driven with plain `docker compose`.
- Docs: updated supported-version references and the compatibility matrix to
  v0.131.1 across `README.md`, `Usage.md`, the CRD descriptions, and the
  consolidated installer (`dist/install.yaml`).

### Removed
- `test/e2e/signoz-override.yaml` — folded into the Foundry casting template's
  `signoz.spec.env`.

### Compatibility
- No operator API changes; the `Alert` and `Endpoint` CRD schemas are unchanged.
- SigNoz-side: the v5 builder-query schema tightened `signal` (logs/metrics/traces)
  and query `type` into **required discriminators** between v0.127 and v0.131.
  Rule bodies targeting v0.131 must set them — the operator forwards verbatim, so
  no caller change. See `Usage.md`.

## [1.127.0] - 2026-06-04

### Changed
- Verified against SigNoz v0.127.0. `ErrorsJSON` gained additive optional fields
  between v0.124 and v0.127; no caller impact.

## [1.124.1] - 2026-05-17

### Fixed
- `kubectl delete alert` no longer hangs after the SigNoz-side rule was removed
  via the UI.

## [1.124.0] - 2026-05-17

### Changed
- First v0.124-line release. The new `/api/v2/infra_monitoring/*` endpoints in
  SigNoz don't affect the operator.

## [1.122.0] - 2026-05-17

### Added
- Initial release.
