package srv

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (ws *Server) Static(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	if page == "/" {
		page = "/index.html"
	}

	ws.ServeFile(w, r, page)
}

func (ws *Server) ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	fn := filepath.Join(ws.staticpath, name)
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
