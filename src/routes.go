package src

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
	"strings"
)

// PopulateRoutes populates the routes of the Moujin instance.
func (m *Moujin) PopulateRoutes() {
	CssDir := http.Dir(filepath.Join(m.Dir, "web", "stylesheets"))
	JsDir := http.Dir(filepath.Join(m.Dir, "web", "scripts"))
	AssDir := http.Dir(filepath.Join(m.Dir, "web", "assets"))
	m.R.Get("/", Home)
	m.R.Get("/home", GetHomePage)
	m.R.Get("/about", GetAboutPage)
	FileServer(m.R, "/css", CssDir)
	FileServer(m.R, "/js", JsDir)
	FileServer(m.R, "/ass", AssDir)
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := HomeTemplate().Render(r.Context(), w)
	if err != nil {
		return
	}
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	err := HomePage().Render(r.Context(), w)
	if err != nil {
		return
	}
}

func GetAboutPage(w http.ResponseWriter, r *http.Request) {
	err := AboutPage().Render(r.Context(), w)
	if err != nil {
		return
	}
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
