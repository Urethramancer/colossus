package main

import (
	"errors"
	"time"

	"github.com/Urethramancer/colossus/internal/osenv"
	"github.com/Urethramancer/colossus/internal/web"
	"github.com/Urethramancer/daemon"
	"github.com/Urethramancer/signor/env"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// CmdServe options.
type CmdServe struct {
	opt.DefaultHelp

	ConfigPath string `short:"C" long:"configpath" help:"Path to configuration files." default:"config"`
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
	Static  string `short:"S" long:"staticpath" help:"Path to static files." default:"static"`
	Domains string `short:"d" long:"domains" help:"Comma-separated list of domains to respond to."`
}

// Run serve
func (cmd *CmdServe) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	ws := &web.Server{}
	ws.DBHost = osenv.Get("DBHOST", cmd.DBHost)
	ws.DBPort = osenv.Get("DBPORT", cmd.DBPort)
	ws.DBName = osenv.Get("DBNAME", cmd.DBName)
	ws.DBUser = osenv.Get("DBUSER", cmd.DBUser)
	ws.DBPass = osenv.Get("DBPASS", cmd.DBPass)
	ws.SSL = env.Get("DB_SSL", cmd.SSL)

	ws.IP = osenv.Get("WEBIP", cmd.IP)
	ws.Port = osenv.Get("WEBPORT", cmd.Port)
	ws.StaticPath = osenv.Get("STATICPATH", cmd.Static)

	ws.Logger = log.Default
	ws.L = log.Default.TMsg
	ws.E = log.Default.TErr

	ws.IdleTimeout = time.Second * 30
	ws.ReadTimeout = time.Second * 10
	ws.WriteTimeout = time.Second * 10

	ws.API = chi.NewRouter()
	ws.API.Use(
		middleware.NoCache,
		web.AddCORS,
		middleware.Timeout(time.Second*10),
	)
	ws.API.NotFound(ws.APInotfound)
	ws.API.Route("/", func(r chi.Router) {
		r.Options("/", web.Preflight)
	})

	ws.Share = chi.NewRouter()
	ws.Share.Use(
		middleware.NoCache,
		middleware.RealIP,
		middleware.RequestID,
		ws.AddLogger,
	)

	ws.Web = chi.NewRouter()
	ws.Web.Use(
		middleware.RealIP,
		middleware.RequestID,
		ws.AddLogger,
		web.AddHTMLHeaders,
	)

	ws.Web.Get("/api", ws.API.ServeHTTP)
	ws.Web.Get("/files", ws.Files)
	ws.Web.Get("/", ws.Static)
	ws.Web.Route("/{page}", func(r chi.Router) {
		r.Get("/*", ws.Static)
	})

	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
	return nil
}
