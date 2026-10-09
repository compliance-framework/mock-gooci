# mock-gooci

Mock repo for developing CCF release automation. Not a product.

It is also the test mock for merging a gh-stack of PRs under strict `ccf-review` approval rules.

This sentence is the second PR of that two-PR stack, based on the first PR's branch.

A miniature of [gooci](https://github.com/compliance-framework/gooci): a Go library
(`pkg/oci`) plus a CLI (`cmd/mock-gooci`) released with goreleaser.

```sh
go test ./...
goreleaser build --snapshot --clean          # linux + darwin binaries in dist/
go install github.com/compliance-framework/mock-gooci/cmd/mock-gooci@<ver>
mock-gooci                                   # prints "mock-gooci <ver>"
```

## Releasing

Releases use the shared [compliance-framework/workflows](https://github.com/compliance-framework/workflows)
release flow (`release-please.yml`, `release-go-lib.yml`, `release-checks.yml`):

1. Merge conventional-commit PRs to `main`; `release-please` opens or updates the release PR
   (`release-please--branches--main`), versioned from `.release-please-manifest.json`
   (the first release is `v0.1.0`).
2. Merging the release PR tags `vX.Y.Z` and publishes the GitHub release.
3. `release` (on `release: published`) runs goreleaser on the tag and attaches the archives
   to that release (`.goreleaser.yaml` keeps `release.prerelease: auto`).
