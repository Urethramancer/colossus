package cfg

import (
	"os"

	"github.com/Urethramancer/colossus/internal/web"
	"github.com/Urethramancer/signor/files"
)

func StartUserWatcher(ws *web.Server, path string) chan bool {
	if !files.DirExists(path) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			ws.E("Error creating '%s': %s", path, err.Error())
			os.Exit(2)
		}

		ws.L("Created %s", path)
	}

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

func StartShareWatcher(ws *web.Server, path string) chan bool {
	if !files.DirExists(path) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			ws.E("Error creating '%s': %s", path, err.Error())
			os.Exit(2)
		}

		ws.L("Created %s", path)
	}

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
