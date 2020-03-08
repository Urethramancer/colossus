package web

import (
	"fmt"
	"net/http"
)

const (
	errUserPassword   = "User and/or password unknown."
	errorInvalidToken = "Invalid token. Please authenticate."
	errorBadPassword  = "Bad or easily guessable password."
)

func (ws *Server) apierror(w http.ResponseWriter, msg string, code int) {
	s := fmt.Sprintf("{\"message\":\"%s\",\"ok\":false}", msg)
	w.Write([]byte(s))
}

func (ws *Server) apinotfound(w http.ResponseWriter, r *http.Request) {
	ws.apierror(w, "Unknown endpoint.", 404)
}

func preflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	w.Header().Set("Access-Control-Max-Age", "86400")
	http.Error(w, "", 204)
}
