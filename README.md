# model-core.go

[![CI](https://github.com/sweetrpg/model-core.go/actions/workflows/ci.yaml/badge.svg)](https://github.com/sweetrpg/model-core.go/actions/workflows/ci.yaml)
[![License](https://img.shields.io/github/license/sweetrpg/model-core.go.svg)](https://img.shields.io/github/license/sweetrpg/model-core.go.svg)
[![Issues](https://img.shields.io/github/issues/sweetrpg/model-core.go.svg)](https://img.shields.io/github/issues/sweetrpg/model-core.go.svg)
[![PRs](https://img.shields.io/github/issues-pr/sweetrpg/model-core.go.svg)](https://img.shields.io/github/issues-pr/sweetrpg/model-core.go.svg)
[![Dependabot](https://badgen.net/github/dependabot/sweetrpg/model-core.go)](https://badgen.net/github/dependabot/sweetrpg/model-core.go)

Base model/value-object types shared across sweetrpg's catalog and data services:
`Auditable`/`AuditableVO` (created/updated/deleted tracking fields), and `Property`/`Tag`
embedded models with conversion helpers between the persistence model and API value object
representations.

## Install

```bash
go get github.com/sweetrpg/model-core.go
```

## Packages

- `models` - `Auditable`, `Property`, `Tag` (persistence-layer structs)
- `vo` - `AuditableVO`, `PropertyVO`, `TagVO` (API-facing equivalents)
- `util` - `FromPropertyModel(s)`/`ToPropertyModel(s)` and `FromTagModel(s)`/`ToTagModel(s)`
  conversion functions between `models` and `vo`

## Documentation

Package documentation: [pkg.go.dev/github.com/sweetrpg/model-core.go](https://pkg.go.dev/github.com/sweetrpg/model-core.go).
Test coverage reports are published to [sweetrpg.github.io/model-core.go](https://sweetrpg.github.io/model-core.go)
on every merge to `develop`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow and
[RELEASE.md](RELEASE.md) for how versions get cut.
