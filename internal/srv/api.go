package srv

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
	println("apinotfound()")
	ws.apierror(w, "Unknown endpoint.", 404)
}

func (ep *Server) apiRootHandler(w http.ResponseWriter, r *http.Request) {
	println("rootHandler()")
	w.Write([]byte("Insert docs here."))
}
