package web

import (
	"net/http"
	"time"
)

// Share or file information.
type Share struct {
	// Name to display for file or directory.
	Name string
	// Path to the file or directory, relative to configured storage path.
	Path string
	// Password is empty if using an account system.
	Password string
	// Created time.
	Created time.Time
	// Users list names with access, and a bool for admin access.
	Users map[string]bool
	// Dir will be set to true if Path points to a directory.
	Dir bool
}

func (ws *Server) Files(w http.ResponseWriter, r *http.Request) {
	ws.wout(w, "No files.")
}
