
## 0.1.0 - 2026-09-02

### Added
- Add SystemActor constant and StampCreate/StampUpdate helpers



## 0.0.173 - 2026-07-23

### Documentation
- Update README (#132)


# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- `models.SystemActor` constant (`"system"`) - the canonical `created_by`/`updated_by` value for
  non-user-driven writes, per the platform audit-fields convention (PADR-0001).
- `models.StampCreate` / `models.StampUpdate` helpers - set the create/update audit fields on an
  `*Auditable` in one call.
- CONTRIBUTING.md, CODE_OF_CONDUCT.md, AGENTS.md/CLAUDE.md repo scaffolding.
- Test coverage for `FromPropertyModel(s)`/`ToPropertyModel(s)` (previously untested - the
  existing property_test.go only exercised the `models.Property` struct fields, never the
  conversion functions themselves).
