// Command mock-gooci prints its version. Release automation installs it with
// `go install github.com/compliance-framework/mock-gooci/cmd/mock-gooci@<ver>`.
package main

import (
	"fmt"
	"runtime/debug"
)

// version is set by goreleaser via -ldflags "-X main.version=...".
var version = ""

// resolveVersion prefers the ldflags value, then the module version recorded
// by `go install ...@<ver>`, and falls back to "dev".
func resolveVersion(ldflags string, info *debug.BuildInfo, ok bool) string {
	if ldflags != "" {
		return ldflags
	}
	if ok && info != nil && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

func main() {
	info, ok := debug.ReadBuildInfo()
	fmt.Println("mock-gooci", resolveVersion(version, info, ok))
}
