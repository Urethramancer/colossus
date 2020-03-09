package cfg

import (
	"os"

	"github.com/fsnotify/fsnotify"

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

	w, err := fsnotify.NewWatcher()
	if err != nil {
		ws.E("Couldn't create watcher': %s", err.Error())
		os.Exit(2)
	}

	err = w.Add(path)
	if err != nil {
		ws.E("Couldn't watch %s: %s", path, err.Error())
		os.Exit(2)
	}

	q := make(chan bool)
	go func() {
		defer w.Close()
		ws.L("User watcher: Start.")
		for {
			select {
			case <-q:
				ws.L("User watcher: Quit.")
				return
			case e := <-w.Events:
				switch {
				case e.Op&fsnotify.Create == fsnotify.Create:
					ws.L("Create %s", e.Name)
				case e.Op&fsnotify.Remove == fsnotify.Remove:
					ws.L("Remove %s", e.Name)
				}
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

	w, err := fsnotify.NewWatcher()
	if err != nil {
		ws.E("Couldn't create watcher': %s", err.Error())
		os.Exit(2)
	}

	err = w.Add(path)
	if err != nil {
		ws.E("Couldn't watch %s: %s", path, err.Error())
		os.Exit(2)
	}

	q := make(chan bool)
	go func() {
		ws.L("Share watcher: Start.")
		for {
			select {
			case <-q:
				ws.L("Share watcher: Quit.")
				return
			case e := <-w.Events:
				switch {
				case e.Op&fsnotify.Create == fsnotify.Create:
					ws.L("Create %s", e.Name)
				case e.Op&fsnotify.Remove == fsnotify.Remove:
					ws.L("Remove %s", e.Name)
				}
			}
		}
	}()
	return q
}
