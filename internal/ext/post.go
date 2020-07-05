// +build allext blogpost

package ext

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Urethramancer/signor/files"
	"github.com/Urethramancer/signor/log"
	"github.com/go-chi/chi"
)

func init() {
	ex := &post{}
	ex.Logger = log.Default
	ex.L = log.Default.TMsg
	ex.E = log.Default.TErr
	RegisterExtension(ex)
	ep := &postapi{}
	ep.Logger = log.Default
	ep.L = log.Default.TMsg
	ep.E = log.Default.TErr
	RegisterEndpoint(ep)
}

type post struct {
	log.LogShortcuts
	cfgpath string
}

func (ex post) Name() string {
	return "post"
}

func (ex *post) LoadConfig(path string) error {
	ex.cfgpath = path
	err := files.EnsureDirExists(path)
	if err != nil {
		return os.ErrExist
	}

	fn := filepath.Join(path, "post.json")
	ex.L("blank: Loading config '%s'.", fn)
	return nil
}

func (ex *post) Pattern() string {
	return "/post"
}

// Routes for blog posts.
func (ex *post) Routes(r chi.Router) {
	r.Get("/{id}", ex.postHandler)
}

func (ex *post) postHandler(w http.ResponseWriter, r *http.Request) {
	wo(w, "Got "+chi.URLParam(r, "id"))
}

//
// API
//

type postapi struct {
	log.LogShortcuts
	cfgpath string
}

func (ep *postapi) Base() string {
	return "/post"
}

// Routes for API GET/POST endpoints.
func (ep *postapi) Routes(r chi.Router) {
	r.Get("/{id}", ep.postHandler)
	// r.Options("/submit", mid.Preflight)
	r.Post("/submit", ep.submitPostHandler)
}

func (ep *postapi) postHandler(w http.ResponseWriter, r *http.Request) {
	println("postHandler()")
	wo(w, "Got "+chi.URLParam(r, "id"))
}

// submitPostHandler accepts a JSON post entry.
// JSON properties:
// datatime - Set to now if unspecified.
// title - Required title for post
// tags - array of tags (categories)
// content - markup containing the entry.
//
// Return JSON
// ok - true if it went well.
// id - the post's ID, if ok is true.
// error - text explaining the problem if ok is false.
func (ep *postapi) submitPostHandler(w http.ResponseWriter, r *http.Request) {
	println("submitPostHandler()")
	wo(w, "New post accepted")
}

func (ep *postapi) idok(w http.ResponseWriter, id string) {
	s := fmt.Sprintf("{\"id\":\"%s\",\"ok\":true}", id)
	w.Write([]byte(s))
}
