package main

import (
	"path/filepath"
	"time"

	"github.com/Urethramancer/colossus/internal/ext"
	"github.com/Urethramancer/colossus/internal/osenv"
	"github.com/Urethramancer/colossus/internal/srv"
	"github.com/Urethramancer/daemon"
	"github.com/Urethramancer/signor/env"
)

func main() {
	datapath := osenv.Get("DATAPATH", "data")
	ws := srv.New(
		osenv.Get("WEBIP", "0.0.0.0"),
		osenv.Get("WEBPORT", "8000"),
		datapath,
		osenv.Get("STATICPATH", "static"),
		osenv.Get("SHAREPATH", "share"),
	)

	ws.SetDatabase(
		osenv.Get("DBHOST", "localhost"),
		osenv.Get("DBPORT", "5432"),
		osenv.Get("DBNAME", "colossus"),
		osenv.Get("DBUSER", "colossus"),
		osenv.Get("DBPASS", "colossus"),
		env.Get("DBSSL", "disabled"),
	)

	ws.IdleTimeout = time.Second * 30
	ws.ReadTimeout = time.Second * 10
	ws.WriteTimeout = time.Second * 10

	list := ext.GetExtensions()
	for k, v := range list {
		ws.WebGets(v.Pattern(), v.Routes)
		ws.L("Extension '%s' added with pattern '%s'.", k, v.Pattern())
		path := filepath.Join(datapath, k)
		v.LoadConfig(path)
	}

	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
}
