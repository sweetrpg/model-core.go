# AGENTS.md

This file provides guidance to Claude Code, Codex, GitHub Copilot, and other AI coding agents
working in this repository.

## About This Project

`model-core.go` provides base model/value-object types shared across sweetrpg's catalog and
data services: `Auditable`/`AuditableVO` (created/updated/deleted tracking fields), and
`Property`/`Tag` embedded models with conversion helpers between the persistence model
(`models`) and API value object (`vo`) representations.

## Dependencies

Depends only on `common.go` (generic slice mapping via `util.Map`). Depended on by
`catalog-objects.go`, `catalog-data.go`, and other sweetrpg data-model packages.

## Committing Code

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

## Branches and Workflow

* `develop` - integration branch, default branch, target for all PRs.
* `master` - latest released state, nothing committed directly.
* `feature/*`, `fix/*` branched from `develop`; `hotfix/*` branched from `master`.

See `CONTRIBUTING.md` for the full workflow.

## Running Checks Locally

```bash
go build -v ./...
go vet ./...
go test -v -coverprofile coverage.out ./...
```

## Releases

See `RELEASE.md`. Summary: trigger `prepare-release.yaml` (`workflow_dispatch` against
`develop`), which computes the next version from conventional commits via git-cliff and opens
a `release/<version>` PR into `master`. Merging that PR tags the release
(`tag-release.yaml`), which triggers `release.yaml` - re-runs tests, creates a GitHub
Release, and merges `master` back into `develop`.
