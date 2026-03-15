# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog.

## [Unreleased]

### Added

- Open-source community health files (`CONTRIBUTING`, `CODE_OF_CONDUCT`, `SECURITY`, `SUPPORT`).
- CI workflow for tests, race detector, and vulnerability scanning.
- Release workflow to publish GitHub releases from tags.
- Issue and pull request templates.
- Package documentation (`doc.go`) and improved README.
- Executable example service in `examples/basic-service`.

### Changed

- `RunAll` now skips nil or invalid runners defensively and logs warnings.
- `RunUntilSignal` now validates runner functions, stops signal notifications, and avoids duplicate stop calls.
- Tests now reserve dynamic ports to reduce CI flakes.
