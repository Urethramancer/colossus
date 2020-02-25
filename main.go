package main

import (
	"os"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

var options struct {
	opt.DefaultHelp
	Init  CmdInit  `command:"init" help:"Initialise database and default admin user."`
	Serve CmdServe `command:"serve" help:"Start the server."`
	User  CmdUser  `command:"user" help:"User management."`
}

func main() {
	a := opt.Parse(&options)
	if options.Help || len(os.Args) < 2 {
		a.Usage()
		return
	}

	err := a.RunCommand(false)
	if err != nil {
		log.Default.Msg("Error running: %s", err.Error())
		os.Exit(2)
	}
}
