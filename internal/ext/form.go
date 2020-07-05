package ext

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Urethramancer/signor/files"
	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
)

func init() {
	ex := &form{}
	ex.Path = "/form"
	ex.Logger = log.Default
	ex.L = log.Default.TMsg
	ex.E = log.Default.TErr
	RegisterExtension(ex)
}

type form struct {
	log.LogShortcuts
	cfgpath string
	formConfig
}

type formConfig struct {
	// Path for form URLs.
	Path string `json:"path"`
}

func (ex form) Name() string {
	return "form"
}

func (ex *form) LoadConfig(path string) error {
	ex.cfgpath = path
	err := files.EnsureDirExists(path)
	if err != nil {
		return os.ErrExist
	}

	fn := filepath.Join(path, "form.json")
	ex.L("form: Loading config '%s'.", fn)
	return nil
}

func (ex *form) Pattern() string {
	return ex.Path
}

// Routes for forms.
func (ex *form) Routes(r chi.Router) {
	r.Get("/{name}", ex.formHandler)
	r.Post("/{name}", ex.formHandler)
}

func (ex *form) formHandler(w http.ResponseWriter, r *http.Request) {
	wo(w, "Got "+chi.URLParam(r, "name"))
}
