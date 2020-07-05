package srv

import (
	"net/http"
)

// AddLogger sets a function to create access log lines.
func (ws *Server) addLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		go func() {
			ws.Logger.TMsg("client %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
