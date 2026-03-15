# Contributing to gootstrap

Thanks for your interest in contributing.

## Development Setup

Requirements:

- Go 1.20+

Run locally:

```bash
go test ./...
```

## Contribution Flow

1. Fork the repository and create a branch from `master`.
2. Keep changes focused and small.
3. Add or update tests for behavior changes.
4. Run `go test ./...` before opening a PR.
5. Open a pull request with context about the problem and the approach.

## Commit Messages

Prefer concise, imperative messages:

- `fix: avoid nil runner panic in RunAll`
- `docs: improve quick start example`

## Pull Request Checklist

- [ ] Tests pass locally
- [ ] New behavior has tests
- [ ] Public API changes are documented in `README.md`
- [ ] Backward compatibility is considered

## Code Style

- Follow idiomatic Go
- Keep API and behavior predictable
- Avoid unnecessary dependencies

## Reporting Issues

When opening an issue, include:

- Go version
- OS
- Repro steps
- Expected vs actual behavior
