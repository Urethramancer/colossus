// Package tpl handles templates for colossus.
// Use from extensions to load arbitrary templates from the right paths.
package tpl

import (
	"path/filepath"
)

var loader struct {
	path      string
	templates []string
}

// NewLoader
func SetPath(path string) {
	loader.path = path
}

// Load a template from a path within the template directory.
// This should be outside the web-servable static paths.
func Load(path string, data interface{}) error {
	fn := filepath.Join(loader.path, path)
	name := filepath.Base(path)
	if name == "" {
		return errNoName{}
	}

	println(fn)
	return nil
}
