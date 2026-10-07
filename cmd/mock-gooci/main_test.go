package main

import (
	"runtime/debug"
	"testing"
)

func TestResolveVersion(t *testing.T) {
	withVersion := func(v string) *debug.BuildInfo {
		return &debug.BuildInfo{Main: debug.Module{Version: v}}
	}
	tests := []struct {
		name    string
		ldflags string
		info    *debug.BuildInfo
		ok      bool
		want    string
	}{
		{"ldflags wins", "1.2.3", withVersion("v9.9.9"), true, "1.2.3"},
		{"go install module version", "", withVersion("v0.1.0"), true, "v0.1.0"},
		{"(devel) placeholder is dev", "", withVersion("(devel)"), true, "dev"},
		{"no build info", "", nil, false, "dev"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveVersion(tt.ldflags, tt.info, tt.ok); got != tt.want {
				t.Errorf("resolveVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}
