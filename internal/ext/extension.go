package ext

import (
	"github.com/go-chi/chi"
)

// Extension defines the interface for new page types.
type Extension interface {
	// Name must be unique.
	Name() string
	// LoadConfig is called with the path to a configuration directory
	// containing the extension's config files.
	LoadConfig(string) error
	// Pattern to hang Routes on.
	Pattern() string
	// Routes has to set up all routes.
	Routes(r chi.Router)
}

var extensions map[string]Extension

// RegisterExtension must be called in an extension's init() function.
func RegisterExtension(ex Extension) error {
	if extensions == nil {
		extensions = make(map[string]Extension)
	}

	// n := filepath.Base(reflect.TypeOf(p).Elem().PkgPath())
	extensions[ex.Name()] = ex
	return nil
}

// GetExtensions returns a map of activated extensions.
func GetExtensions() map[string]Extension {
	return extensions
}
