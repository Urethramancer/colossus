package web

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	fi, err := os.Stat(path)
	if err != nil {
		ws.E("Error sharing '%s': %s", name, err.Error())
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ws.E("Error sharing:'%s': %s", err.Error())
		return
	}

	sh := &Share{
		Name:     name,
		Path:     path,
		Password: string(pw),
		Created:  time.Now(),
		Users:    make(map[string]bool),
		Dir:      fi.IsDir(),
	}

	ws.shares[name] = sh
}

// RemoveShare completely removes a share and its configuration file.
func (ws *Server) RemoveShare(name string) {
	sh, ok := ws.shares[name]
	if !ok {
		ws.E("Unknown share '%s.", name)
		return
	}

	err := os.RemoveAll(sh.Path)
	if err != nil {
		ws.E("Error deleting '%s': %s", sh.Path, err.Error())
	}

	delete(ws.shares, name)
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
