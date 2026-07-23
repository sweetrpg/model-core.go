# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- CONTRIBUTING.md, CODE_OF_CONDUCT.md, AGENTS.md/CLAUDE.md repo scaffolding.
- Test coverage for `FromPropertyModel(s)`/`ToPropertyModel(s)` (previously untested - the
  existing property_test.go only exercised the `models.Property` struct fields, never the
  conversion functions themselves).
