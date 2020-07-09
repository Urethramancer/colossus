package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Urethramancer/colossus/accounts"
	"github.com/Urethramancer/colossus/internal/ext"
	"github.com/Urethramancer/colossus/internal/srv"
	"github.com/Urethramancer/daemon"
)

func get(k string) func(*srv.Server) {
	return func(srv *srv.Server) {
		v := os.Getenv(k)
		if v != "" {
			srv.Set(k, v)
		}
	}
}

func getm(k string) func(*accounts.Manager) {
	return func(m *accounts.Manager) {
		v := os.Getenv(k)
		if v != "" {
			m.Set(k, v)
		}
	}
}

func main() {
	ws := srv.New(
		get(srv.ENVHOST),
		get(srv.ENVPORT),
		get(srv.ENVDATA),
		get(srv.ENVSTATIC),
		get(srv.ENVSHARE),
	)

	ws.IdleTimeout = time.Second * 30
	ws.ReadTimeout = time.Second * 10
	ws.WriteTimeout = time.Second * 10

	list := ext.GetExtensions()
	for k, v := range list {
		ws.WebGets(v.Pattern(), v.Routes)
		ws.L("Extension '%s' added with pattern '%s'.", k, v.Pattern())
		path := filepath.Join(os.Getenv(srv.ENVDATA), k)
		v.LoadConfig(path)
	}

	m, err := accounts.NewManager(
		getm(accounts.ENVDBHOST),
		getm(accounts.ENVDBPORT),
		getm(accounts.ENVDBNAME),
		getm(accounts.ENVDBUSER),
		getm(accounts.ENVDBPASS),
		getm(accounts.ENVDBSSL),
	)
	if err != nil {
		ws.E("Error setting up account manager: %s", err.Error())
		os.Exit(2)
	}

	users, err := m.GetUsers(10)
	if err == nil && len(users) == 0 {
		pw, err := m.CreateRootUser("admin", os.Getenv(accounts.ENVEMAIL))
		if err != nil {
			ws.E("Couldn't create initial admin user: %s", err.Error())
			os.Exit(2)
		}

		ws.L("Superuser admin created with password %s", pw)
	}

	defer m.Close()
	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
}
