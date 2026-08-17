# lotof.sample.proto

Template `*.proto` package for bootstrapping a new LOTOF domain. Published as
a Go module (`github.com/pieceowater-dev/lotof.sample.proto`) and consumed by
[lotof.sample.svc](https://github.com/pieceowater-dev/lotof.sample.svc) and
[lotof.sample.gtw](https://github.com/pieceowater-dev/lotof.sample.gtw), the
matching microservice and gateway templates.

## Layout

```
protos/
  generic/
    utils/
      filter.proto       # Pagination (list request input)
      list.proto         # PaginationInfo (list response output)
      middleware.proto   # UserMetadata — per-request identity/tenant context
    tenants/
      tenants.proto       # GatewayService (Hub->this app) + AppTenantsService (this app's own tenant provisioning, implemented by the svc)
  lotof.sample.svc/
    domainItem/
      domainItem.proto     # Example CRUD service — delete/rename when bootstrapping a real domain
```

`generic/utils` and `generic/tenants` are the multi-tenant scaffolding shared
by every current domain (menu, issues, contacts, atrace) — keep them as-is.
`lotof.sample.svc/domainItem` is the disposable example; see the svc/gtw
templates' READMEs for what to rename.

## Bootstrapping a new domain

1. Fork this repo as `lotof.<domain>.proto`.
2. Update the module path in `go.mod` to `github.com/pieceowater-dev/lotof.<domain>.proto`.
3. Rename `protos/lotof.sample.svc/domainItem/` to `protos/lotof.<domain>.msvc.<name>/<yourEntity>/` and edit the service/message names.
4. Keep `protos/generic/` untouched — it's shared multi-tenant plumbing, not domain-specific.
5. Push to `main` — CI (`.github/workflows/release.yml`) auto-tags a new `v0.0.X` patch version and publishes the Go module.

## Publishing

Every push to `main` auto-increments the patch version tag (`v0.0.X` ->
`v0.0.X+1`) and pushes it, which triggers GOPROXY indexing so consumers can
immediately `go get github.com/pieceowater-dev/lotof.<domain>.proto@v0.0.X+1`.
Private module access from consuming repos' CI requires the `GH_PRIVATE_PKG_PAT`
secret.
