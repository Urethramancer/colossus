package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

// CmdUser options.
type CmdUser struct {
	opt.DefaultHelp
	Add CmdUserAdd `command:"add" help:"Add a new user."`
}

// Run user commands
func (cmd *CmdUser) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

// CmdUserAdd options.
type CmdUserAdd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:""`
	Path string `placeholder:"PATH" help:""`
}

// Run user add
func (cmd *CmdUserAdd) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}
