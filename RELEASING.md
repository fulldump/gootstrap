# Releasing

## Create a release

1. Ensure `master` is green in CI.
2. Update `CHANGELOG.md`.
3. Create and push a semantic version tag.

```bash
git tag v1.0.0
git push origin v1.0.0
```

Pushing a `v*` tag triggers `.github/workflows/release.yml` and creates a
GitHub release with generated notes.
