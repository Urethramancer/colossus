package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

// CmdUser options.
type CmdUser struct {
	opt.DefaultHelp
	Add    CmdUserAdd    `command:"add" help:"Add a new user." aliases:"a"`
	Remove CmdUserRemove `command:"remove" help:"Remove a user." aliases:"rm"`
	Reset  CmdUserReset  `command:"reset" help:"Add a new user." aliases:"r"`
	List   CmdUserList   `command:"list" help:"List users for a path." aliases:"ls"`
}

// Run user commands
func (cmd *CmdUser) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

// CmdUserAdd options.
type CmdUserAdd struct {
	opt.DefaultHelp
	Name     string `placeholder:"NAME" help:"Username to create."`
	Path     string `placeholder:"PATH" help:"File directory this user is associated with."`
	Generate bool   `short:"g" long:"generate" help:"Generate a password rather than ask for one."`
}

// Run user add
func (cmd *CmdUserAdd) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdUserRemove options.
type CmdUserRemove struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Username to remove."`
	Path string `placeholder:"PATH" help:"File directory this user is associated with."`
}

// Run user remove
func (cmd *CmdUserRemove) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdUserReset options.
type CmdUserReset struct {
	opt.DefaultHelp
	Name     string `placeholder:"NAME" help:"Username to reset the password for."`
	Path     string `placeholder:"PATH" help:"File directory this user is associated with."`
	Generate bool   `short:"g" long:"generate" help:"Generate a password rather than ask for one."`
}

// Run user reset
func (cmd *CmdUserReset) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdUserList options.
type CmdUserList struct {
	opt.DefaultHelp
	Path string `placeholder:"PATH" help:"File directory to list users from."`
}

// Run user list
func (cmd *CmdUserList) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}
