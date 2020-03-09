package main

import (
	"errors"
	"time"

	"github.com/Urethramancer/colossus/internal/osenv"
	"github.com/Urethramancer/daemon"
	"github.com/Urethramancer/signor/env"
	"github.com/Urethramancer/signor/opt"
)

// CmdServe options.
type CmdServe struct {
	opt.DefaultHelp

	DataPath string `short:"D" long:"configpath" help:"Path to user/share configuration files." default:"data"`
	// Database options
	DBHost string `short:"H" long:"dbhost" help:"Database host to connect to." default:"localhost"`
	DBPort string `short:"p" long:"dbport" help:"Database port to connect to." default:"5432"`
	DBName string `short:"n" long:"dbname" help:"Database to open." default:"colossus"`
	DBUser string `short:"u" long:"dbuser" help:"Database user to connect as." default:"postgres"`
	DBPass string `short:"P" long:"dbpass" help:"Database password to connect with. Nay be blank."`
	SSL    string `short:"s" long:"ssl" help:"Require SSL to connect to the database." choices:"enable,disable" default:"disable"`

	// Webserver options
	IP     string `short:"i" long:"ip" help:"IP address to bind to." default:"127.0.0.1"`
	Port   string `short:"w" long:"port" help:"Port to run on." default:"8000"`
	Static string `short:"S" long:"staticpath" help:"Path to static files." default:"static"`
	Shared string `short:"F" long:"sharepath" help:"Path to store shared files and folders." default:"shared"`
	// Domains string `short:"d" long:"domains" help:"Comma-separated list of domains to respond to."`
}

// Run serve
func (cmd *CmdServe) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	ws := NewServer(
		osenv.Get("WEBIP", cmd.IP),
		osenv.Get("WEBPORT", cmd.Port),
		osenv.Get("DATAPATH", cmd.DataPath),
		osenv.Get("STATICPATH", cmd.Static),
		osenv.Get("SHAREPATH", cmd.Shared),
	)

	ws.SetDatabase(
		osenv.Get("DBHOST", cmd.DBHost),
		osenv.Get("DBPORT", cmd.DBPort),
		osenv.Get("DBNAME", cmd.DBName),
		osenv.Get("DBUSER", cmd.DBUser),
		osenv.Get("DBPASS", cmd.DBPass),
		env.Get("DBSSL", cmd.SSL),
	)

	ws.IdleTimeout = time.Second * 30
	ws.ReadTimeout = time.Second * 10
	ws.WriteTimeout = time.Second * 10

	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
	return nil
}
