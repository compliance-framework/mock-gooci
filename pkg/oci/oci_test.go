package oci

import "testing"

func TestRef(t *testing.T) {
	tests := []struct {
		name, repo, tag, want string
	}{
		{"with tag", "ghcr.io/compliance-framework/plugin", "v1.2.3", "ghcr.io/compliance-framework/plugin:v1.2.3"},
		{"empty tag defaults to latest", "ghcr.io/compliance-framework/plugin", "", "ghcr.io/compliance-framework/plugin:latest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ref(tt.repo, tt.tag); got != tt.want {
				t.Errorf("Ref(%q, %q) = %q, want %q", tt.repo, tt.tag, got, tt.want)
			}
		})
	}
}
