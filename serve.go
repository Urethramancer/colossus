package main

import (
	"errors"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"

	"github.com/Urethramancer/colossus/osenv"
	"github.com/Urethramancer/daemon"
	"github.com/Urethramancer/signor/env"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

// CmdServe options.
type CmdServe struct {
	opt.DefaultHelp

	// Database options
	DBHost string `short:"H" long:"dbhost" help:"Database host to connect to." default:"localhost"`
	DBPort string `short:"p" long:"dbport" help:"Database port to connect to." default:"5432"`
	DBName string `short:"n" long:"dbname" help:"Database to open." default:"colossus"`
	DBUser string `short:"u" long:"dbuser" help:"Database user to connect as." default:"postgres"`
	DBPass string `short:"P" long:"dbpass" help:"Database password to connect with. Nay be blank."`
	SSL    string `short:"s" long:"ssl" help:"Require SSL to connect to the database." choices:"enable,disable" default:"disable"`

	// Webserver options
	IP      string `short:"i" long:"ip" help:"IP address to bind to." default:"127.0.0.1"`
	Port    string `short:"w" long:"port" help:"Port to run on." default:"8000"`
	Domains string `short:"d" long:"domains" help:"Comma-separated list of domains to respond to."`
}

// Run serve
func (cmd *CmdServe) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	ws := &WebServer{}
	ws.dbhost = osenv.Get("DBHOST", cmd.DBHost)
	ws.dbport = osenv.Get("DBPORT", cmd.DBPort)
	ws.dbname = osenv.Get("DBNAME", cmd.DBName)
	ws.dbuser = osenv.Get("DBUSER", cmd.DBUser)
	ws.dbpass = osenv.Get("DBPASS", cmd.DBPass)
	ssl := env.Get("DB_SSL", cmd.SSL)
	if ssl == "" {
		ssl = cmd.SSL
	}

	ws.ip = osenv.Get("WEBIP", cmd.IP)
	ws.port = osenv.Get("WEBPORT", cmd.Port)

	ws.L = log.Default.TMsg
	ws.E = log.Default.TErr

	ws.IdleTimeout = time.Second * 30
	ws.ReadTimeout = time.Second * 10
	ws.WriteTimeout = time.Second * 10

	ws.api = chi.NewRouter()
	ws.api.Use(middleware.NoCache)
	ws.api.Use(addCORS)
	ws.api.Use(middleware.RealIP)
	ws.api.Use(middleware.RequestID)
	ws.api.Use(middleware.Timeout(time.Second * 10))
	ws.api.NotFound(notfound)
	ws.api.Route("/", func(r chi.Router) {
		r.Options("/", preflight)
	})

	ws.web = chi.NewRouter()
	ws.web.Get("/", static)
	ws.web.Get("/api", ws.api.ServeHTTP)

	m := log.Default.Msg
	m("%+v", ws)

	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
	return nil
}
