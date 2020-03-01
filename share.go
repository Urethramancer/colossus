package main

import "net/http"

func (ws *WebServer) files(w http.ResponseWriter, r *http.Request) {
	ws.wout(w, "No files.")
}
