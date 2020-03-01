package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
)

// WebServer main structure.
type WebServer struct {
	sync.RWMutex
	sync.WaitGroup
	log.LogShortcuts
	http.Server

	dbhost string
	dbport string
	dbname string
	dbuser string
	dbpass string
	ssl    string
	// db     *anthropoi.DBM

	ip         string
	port       string
	staticpath string

	// api endpoints
	api *chi.Mux
	// web server root path
	web *chi.Mux
	// share folders and files
	share *chi.Mux
}

// Start serving.
func (ws *WebServer) Start() {
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
		ws.Handler = ws.web
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
func (ws *WebServer) Stop() {
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

func (ws *WebServer) wout(w http.ResponseWriter, s string) {
	n, err := w.Write([]byte(s))
	if err != nil {
		ws.Logger.TErr("Error: wrote %d bytes: %s", n, err.Error())
	}
}
