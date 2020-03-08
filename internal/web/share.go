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

// AddShare creates a file or directory share, wiht optional global password.
func (ws *Server) AddShare(name, path, password string) {

}

// RemoveShare completely removes a share and its configuration file.
func (ws *Server) RemoveShare(name string) {

}

// AddShareUser adds a user to a share, optionally with admin access.
func (ws *Server) AddShareUsers(name string, admin bool) {
	sh, ok := ws.shares[name]
	if !ok {
		ws.E("Unknown share '%s.", name)
		return
	}

	sh.Users[name] = admin
}
