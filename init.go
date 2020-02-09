package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

// CmdInit options.
type CmdInit struct {
	opt.DefaultHelp
	Drop bool `short:"D" long:"drop" help:"Drop existing database or tables. Requires superuser access."`
}

// Run init
func (cmd *CmdInit) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}
