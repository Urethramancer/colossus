// +build allext blank

package ext

// This is an example extension.

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Urethramancer/signor/files"
	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
)

func init() {
	ex := &blank{}
	ex.Logger = log.Default
	ex.L = log.Default.TMsg
	ex.E = log.Default.TErr
	RegisterExtension(ex)
}

type blank struct {
	log.LogShortcuts
	cfgpath string
}

func (ex blank) Name() string {
	return "blank"
}

func (ex *blank) LoadConfig(path string) error {
	ex.cfgpath = path
	err := files.EnsureDirExists(path)
	if err != nil {
		return os.ErrExist
	}

	fn := filepath.Join(path, "blank.json")
	ex.L("blank: Loading config '%s'.", fn)
	return nil
}

func (ex *blank) Pattern() string {
	return "/blank"
}

// Routes for blank pages.
func (ex *blank) Routes(r chi.Router) {
	r.Get("/{id}", ex.blankHandler)
}

func (ex *blank) blankHandler(w http.ResponseWriter, r *http.Request) {
	wo(w, "Got "+chi.URLParam(r, "id"))
}
