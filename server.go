package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server main structure.
type Server struct {
	sync.RWMutex
	sync.WaitGroup
	log.LogShortcuts
	http.Server

	// address
	dbhost string
	// port number
	dbport string
	// name to connect to
	dbname string
	// user to connect as
	dbuser string
	// password to authenticate with
	dbpass string
	// SSL enabled or disabled
	ssl string
	// db     *anthropoi.DBM

	ip       string
	port     string
	datapath string
	// staticpath is for files retrieved by the client (HTML, CSS, images, JS).
	staticpath string
	sharepath  string

	// api endpoints
	api *chi.Mux
	// web server root path
	web *chi.Mux
	// share folders and files
	share *chi.Mux

	// users are loaded into this
	users map[string]*User
	// shares are loaded into this
	shares map[string]*Share

	userquit  chan bool
	sharequit chan bool
}

// New web server strcture is returned with the essentials filled in.
func NewServer(addr, p, dp, static, shares string) *Server {
	ws := &Server{
		ip:         addr,
		port:       p,
		datapath:   dp,
		staticpath: static,
		sharepath:  shares,
	}

	// Logging
	ws.Logger = log.Default
	ws.L = log.Default.TMsg
	ws.E = log.Default.TErr

	// API middleware and preflight
	ws.api = chi.NewRouter()
	ws.api.Use(
		middleware.NoCache,
		addCORS,
		middleware.RealIP,
		middleware.Timeout(time.Second*10),
	)
	ws.api.NotFound(ws.apinotfound)
	ws.api.Route("/", func(r chi.Router) {
		r.Options("/", preflight)
	})

	// File share setup
	ws.share = chi.NewRouter()
	ws.share.Use(
		middleware.NoCache,
		middleware.RealIP,
		middleware.RequestID,
		ws.addLogger,
	)

	// HTTP middleware
	ws.web = chi.NewRouter()
	ws.web.Use(
		middleware.RealIP,
		middleware.RequestID,
		ws.addLogger,
		addHTMLHeaders,
	)

	ws.WebGet("/api", ws.api.ServeHTTP)
	ws.WebGet("/files", ws.share.ServeHTTP)
	ws.share.Route("/", func(r chi.Router) {
		r.Get("/*", ws.Files)
		ws.L("Added placeholder page for /files root page.")
	})

	ws.WebGet("/", ws.Static)
	ws.WebGets("/{page}", func(r chi.Router) {
		r.Get("/*", ws.Static)
	})

	return ws
}

// Start serving.
func (ws *Server) Start() {
	ws.Lock()
	defer ws.Unlock()

	addr := net.JoinHostPort(ws.ip, ws.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		ws.E("Listener error: %s", err.Error())
		// ws.db.Close()
		os.Exit(2)
	}

	ws.Add(1)
	ws.L("Starting web server on http://%s", addr)
	go func() {
		path := filepath.Join(ws.datapath, "users")
		ws.startUserWatcher(path)
		path = filepath.Join(ws.datapath, "shares")
		ws.startShareWatcher(path)

		ws.Handler = ws.web
		err = ws.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			ws.E("Error running server: %s", err.Error())
			// ws.db.Close()
			ws.userquit <- true
			ws.sharequit <- true
			os.Exit(2)
		}
		ws.L("Stopped web server.")
		ws.Done()
	}()
}

// Stop serving.
func (ws *Server) Stop() {
	ws.userquit <- true
	ws.sharequit <- true
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
