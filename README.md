# mock-gooci

Mock repo for developing CCF release automation. Not a product.

A miniature of [gooci](https://github.com/compliance-framework/gooci): a Go library
(`pkg/oci`) plus a CLI (`cmd/mock-gooci`) released with goreleaser.

```sh
go test ./...
goreleaser build --snapshot --clean          # linux + darwin binaries in dist/
go install github.com/compliance-framework/mock-gooci/cmd/mock-gooci@<ver>
mock-gooci                                   # prints "mock-gooci <ver>"
```
