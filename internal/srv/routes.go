package srv

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (ws *Server) APIGet() {

}

// WebGet adds a GET route matching the specified pattern.
func (ws *Server) WebGet(pattern string, handler http.HandlerFunc) {
	ws.web.Get(pattern, handler)
	ws.L("Added GET route for %s", pattern)
}

// WebGets adds one or more GET routes to the specified pattern.
func (ws *Server) WebGets(pattern string, fn func(r chi.Router)) {
	ws.web.Route(pattern, fn)
	ws.L("Added GET routes for %s", pattern)
}
