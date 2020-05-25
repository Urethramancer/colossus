package main

import (
	"os"
	"time"

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

	// list := ext.GetExtensions()

	// for k, v := range list {
	// 	ws.WebGets(v.Pattern(), v.Routes)
	// 	ws.L("Extension '%s' added with pattern '%s'.", k, v.Pattern())
	// 	path := filepath.Join(datapath, k)
	// 	v.LoadConfig(path)
	// }

	ws.Start()
	<-daemon.BreakChannel()
	ws.Stop()
}
