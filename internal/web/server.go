package web

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Urethramancer/colossus/internal/acc"
	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
)

// Server main structure.
type Server struct {
	sync.RWMutex
	sync.WaitGroup
	log.LogShortcuts
	http.Server

	// DBHost address.
	DBHost string
	// DBPort number.
	DBPort string
	// DBName to connect to.
	DBName string
	// DBUser to connect as.
	DBUser string
	// DBPass to authenticate with.
	DBPass string
	// SSL enabled or disabled.
	SSL string
	// db     *anthropoi.DBM

	IP   string
	Port string
	// staticpath is for files retrieved by the client (HTML, CSS, images, JS).
	StaticPath string

	// api endpoints
	API *chi.Mux
	// web server root path
	Web *chi.Mux
	// share folders and files
	Share *chi.Mux

	// users are loaded into this
	users map[string]*acc.User
	// shares are loaded into this
	shares map[string]*Share
}

// Start serving.
func (ws *Server) Start() {
	ws.Lock()
	defer ws.Unlock()

	addr := net.JoinHostPort(ws.IP, ws.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		ws.E("Listener error: %s", err.Error())
		// ws.db.Close()
		os.Exit(2)
	}

	ws.Add(1)
	ws.L("Starting web server on http://%s", addr)
	go func() {
		ws.Handler = ws.Web
		err = ws.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			ws.E("Error running server: %s", err.Error())
			// ws.db.Close()
			os.Exit(2)
		}
		ws.L("Stopped web server.")
		ws.Done()
	}()
}

// Stop serving.
func (ws *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := ws.Shutdown(ctx)
	if err != nil {
		ws.E("Shutdown error: %s", err.Error())
		os.Exit(2)
	}

	ws.Wait()
	// ws.db.Close()
}

func (ws *Server) wout(w http.ResponseWriter, s string) {
	n, err := w.Write([]byte(s))
	if err != nil {
		ws.E("Error: wrote %d bytes: %s", n, err.Error())
	}
}
