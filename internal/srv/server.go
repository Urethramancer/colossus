package srv

import (
	"context"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Urethramancer/colossus/internal/ext"
	"github.com/Urethramancer/colossus/internal/settings"
	"github.com/Urethramancer/colossus/mid"
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

	settings.Settings

	IP       string
	Port     string
	Datapath string
	// staticpath is for files retrieved by the client (HTML, CSS, images, JS).
	StaticPath string
	SharePath  string

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
// func New(addr, p, dp, static, shares string) *Server {
func New(options ...func(*Server)) *Server {
	ws := &Server{}

	ws.InitVars(map[string]string{
		ENVHOST:   "0.0.0.0",
		ENVPORT:   "8000",
		ENVDATA:   "data",
		ENVSTATIC: "static",
		ENVSHARE:  "share",
	})

	for _, o := range options {
		o(ws)
	}

	// Logging
	ws.Logger = log.Default
	ws.L = log.Default.TMsg
	ws.E = log.Default.TErr

	// HTTP middleware and routing (required by everything else)
	ws.web = chi.NewRouter()
	ws.web.Use(
		middleware.RealIP,
		middleware.RequestID,
		ws.addLogger,
		mid.AddHTMLHeaders,
	)

	// File share setup
	ws.share = chi.NewRouter()
	ws.share.Use(
		middleware.NoCache,
		middleware.RealIP,
		middleware.RequestID,
		ws.addLogger,
	)

	ws.WebGet("/files", ws.share.ServeHTTP)
	ws.share.Route("/", func(r chi.Router) {
		r.Get("/*", ws.Files)
		ws.L("Added placeholder page for /files root page.")
	})

	// API setup from extensions
	// ws.WebGet("/api", ws.api.ServeHTTP)
	ws.web.Route("/api", func(r chi.Router) {
		list := ext.GetEndpoints()
		// API middleware and preflight
		r.Use(
			middleware.NoCache,
			middleware.RealIP,
			mid.AddCORS,
			middleware.Timeout(time.Second*10),
		)
		r.NotFound(ws.apinotfound)
		r.Options("/", mid.Preflight)
		r.Get("/", ws.apiRootHandler)
		for _, ep := range list {
			r.Route(ep.Base(), ep.Routes)
			ws.L("Added /api%s routes.", ep.Base())
		}
	})

	// Static pages (shortest route added last)
	ws.WebGet("/", ws.Static)
	ws.WebGets("/{page}", func(r chi.Router) {
		r.Get("/*", ws.Static)
		r.Options("/", mid.Preflight)
	})

	return ws
}

// Start serving.
func (ws *Server) Start() {
	ws.Lock()
	defer ws.Unlock()

	addr := net.JoinHostPort(ws.Get(ENVHOST), ws.Get(ENVPORT))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		ws.E("Listener error: %s", err.Error())
		// ws.db.Close()
		os.Exit(2)
	}

	ws.Add(1)
	ws.L("Starting web server on http://%s", addr)
	go func() {
		path := filepath.Join(ws.Get(ENVDATA), "users")
		ws.startUserWatcher(path)
		path = filepath.Join(ws.Get(ENVSHARE), "share")
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
