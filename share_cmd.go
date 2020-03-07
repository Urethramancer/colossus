package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

// CmdShare options.
type CmdShare struct {
	opt.DefaultHelp
	Add    CmdShareAdd    `command:"add" help:"Add a new share." aliases:"a"`
	Remove CmdShareRemove `command:"remove" help:"Remove a share." aliases:"rm"`
	Reset  CmdShareReset  `command:"reset" help:"Add a new share." aliases:"r"`
	List   CmdShareList   `command:"list" help:"List shares for a path." aliases:"ls"`
}

// Run share commands
func (cmd *CmdShare) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

// CmdShareAdd options.
type CmdShareAdd struct {
	opt.DefaultHelp
	Name     string `placeholder:"NAME" help:"Identifier of share to create."`
	Path     string `placeholder:"PATH" help:"File directory this share is associated with."`
	Generate bool   `short:"g" long:"generate" help:"Generate a password rather than ask for one."`
}

// Run share add
func (cmd *CmdShareAdd) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdShareRemove options.
type CmdShareRemove struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"sharename to remove."`
	Path string `placeholder:"PATH" help:"File directory this share is associated with."`
}

// Run share remove
func (cmd *CmdShareRemove) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdShareReset options.
type CmdShareReset struct {
	opt.DefaultHelp
	Name     string `placeholder:"NAME" help:"sharename to reset the password for."`
	Path     string `placeholder:"PATH" help:"File directory this share is associated with."`
	Generate bool   `short:"g" long:"generate" help:"Generate a password rather than ask for one."`
}

// Run share reset
func (cmd *CmdShareReset) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

// CmdShareList options.
type CmdShareList struct {
	opt.DefaultHelp
	Path string `placeholder:"PATH" help:"File directory to list shares from."`
}

// Run share list
func (cmd *CmdShareList) Run(in []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}
