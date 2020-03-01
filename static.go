package main

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (ws *WebServer) static(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	if page == "/" {
		page = "/index.html"
	}

	ws.servefile(w, r, page)
}

func (ws *WebServer) servefile(w http.ResponseWriter, r *http.Request, name string) {
	ws.L("%s", name)
	fn := filepath.Join(ws.staticpath, name)
	ws.L("%s", fn)
	f, err := os.Open(fn)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	defer f.Close()
	ext := filepath.Ext(fn)
	if ext != "" {
		w.Header().Set("Content-Type", mime.TypeByExtension(ext))
	} else {
		w.Header().Set("Content-Type", mime.TypeByExtension(".txt"))
	}

	http.ServeContent(w, r, name, time.Time{}, f)
}
