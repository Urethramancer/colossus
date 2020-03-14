// +build allext blank

package ext

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func init() {
	ex := &blank{}
	RegisterExtension(ex)
}

type blank struct {
}

func (ex blank) Name() string {
	return "blank"
}

func (ex *blank) Pattern() string {
	return "/blank"
}

// Routes for blank pages.
func (ex *blank) Routes(r chi.Router) {
	r.Get("/{id}", ex.blankHandler)
}

func (ex *blank) blankHandler(w http.ResponseWriter, r *http.Request) {
	ex.o(w, "Got "+chi.URLParam(r, "id"))
}

func (ex *blank) o(w http.ResponseWriter, s string) {
	n, err := w.Write([]byte(s))
	if err != nil {
		fmt.Printf("Error: wrote %d bytes: %s\n", n, err.Error())
	}
}
