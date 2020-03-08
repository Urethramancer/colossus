package cfg

import "github.com/Urethramancer/colossus/internal/web"

func StartUserWatcher(ws *web.Server) chan bool {
	q := make(chan bool)
	go func() {
		ws.L("User watcher: Start.")
		for {
			select {
			case <-q:
				ws.L("User watcher: Quit.")
				return
			}
		}
	}()
	return q
}

func StartShareWatcher(ws *web.Server) chan bool {
	q := make(chan bool)
	go func() {
		ws.L("Share watcher: Start.")
		for {
			select {
			case <-q:
				ws.L("Share watcher: Quit.")
				return
			}
		}
	}()
	return q
}
