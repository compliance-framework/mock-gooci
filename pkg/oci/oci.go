// Package oci is a stand-in for gooci's OCI helpers. It exists so that
// downstream mock repos have an importable library to pin and bump.
package oci

// Ref joins an OCI repository name and a tag into a reference of the form
// "name:tag". An empty tag defaults to "latest".
func Ref(name, tag string) string {
	if tag == "" {
		tag = "latest"
	}
	return name + ":" + tag
}
