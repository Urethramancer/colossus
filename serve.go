package main

import (
	"errors"

	"github.com/Urethramancer/colossus/osenv"
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

	dbhost := osenv.Get("DBHOST", cmd.DBHost)
	dbport := osenv.Get("DBPORT", cmd.DBPort)
	dbname := osenv.Get("DBNAME", cmd.DBName)
	dbuser := osenv.Get("DBUSER", cmd.DBUser)
	dbpass := osenv.Get("DBPASS", cmd.DBPass)
	ssl := env.Get("DB_SSL", cmd.SSL)
	if ssl == "" {
		ssl = cmd.SSL
	}

	ip := osenv.Get("WEBIP", cmd.IP)
	port := osenv.Get("WEBPORT", cmd.Port)

	m := log.Default.Msg
	m("DBHost: %s\nDBPort: %s\nDBName: %s\nSSL: %s\nDBUser: %s\nDBPass: %s\nIP: %s\nPort: %s\n", dbhost, dbport, dbname, ssl, dbuser, dbpass, ip, port)
	return nil
}
