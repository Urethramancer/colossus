package ext

import (
	"net/http"

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

// Endpoint registration interface.
type Endpoint interface {
	Base() string
	Routes(r chi.Router)
}

var extensions map[string]Extension
var endpoints map[string]Endpoint

// RegisterExtension must be called in an extension's init() function.
func RegisterExtension(ex Extension) {
	if extensions == nil {
		extensions = make(map[string]Extension)
	}

	// n := filepath.Base(reflect.TypeOf(p).Elem().PkgPath())
	extensions[ex.Name()] = ex
}

// GetExtensions returns a map of activated extensions.
func GetExtensions() map[string]Extension {
	return extensions
}

func wo(w http.ResponseWriter, s string) {
	w.Write([]byte(s))
}

// RegisterEndpoint registers a new API endpoint handled by an extension.
func RegisterEndpoint(ep Endpoint) {
	if endpoints == nil {
		endpoints = make(map[string]Endpoint)
	}

	endpoints[ep.Base()] = ep
}

// GetEndpoints returns a map of extension endpoints.
func GetEndpoints() map[string]Endpoint {
	return endpoints
}
