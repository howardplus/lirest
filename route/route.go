package route

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/howardplus/lirest/describe"
	"github.com/urfave/negroni"
	"net/http"
)

// NewRouter
// the router is defined hierarchically
// starting from the top level
func NewRouter(t *Trie) http.Handler {

	// r contains the real routes
	r := mux.NewRouter()

	// create command routes
	GenerateCommandRoutes(r)

	// now walk the trie and create the routes
	generateRoutes(r, t)

	// after routers are created, create a dummy world router
	// world router is simply used to wrap the negroni middleware
	world := mux.NewRouter()

	// tag is available on GET only
	world.Methods("GET").Handler(
		negroni.New(
			negroni.HandlerFunc(LogMiddleware),
			negroni.HandlerFunc(TagMiddleware),
			negroni.HandlerFunc(JsonpMiddleware),
			negroni.Wrap(r),
		))

	// for everything else
	world.Methods("POST", "PUT", "DELETE").Handler(
		negroni.New(
			negroni.HandlerFunc(LogMiddleware),
			negroni.HandlerFunc(JsonpMiddleware),
			negroni.Wrap(r),
		))

	return world
}

// walk through the route trie
// and populate the mux router
func generateRoutes(r *mux.Router, root *Trie) error {
	traverseDepth(r, root, []string{})
	return nil
}

func traverseDepth(r *mux.Router, t *Trie, path []string) {

	var fullpath bytes.Buffer
	fullpath.WriteString("/")
	for _, p := range path {
		fullpath.WriteString(p + "/")
	}

	// install handlers
	if t.Val != nil {
		desc := t.Val.(describe.Description)

		log.WithFields(log.Fields{
			"name":    desc.Name,
			"methods": desc.Api.Methods,
			"path":    desc.Api.Path,
		}).Debug("Route description")

		for _, method := range desc.Api.Methods {

			log.WithFields(log.Fields{
				"name":     desc.Name,
				"method":   method,
				"path":     fullpath.String(),
				"api-path": desc.Api.Path,
			}).Debug("Install resource handler")

			// create resource handlers
			r.Methods(method).Path(desc.Api.Path).Handler(
				&ResourceHandler{
					desc.Name,
					desc.System,
					desc.Api,
					desc.Vars,
				})
		}
	} else {
		// find all the possible subpath from here onwards
		subpath := []string{}
		for k := range t.Nodes {
			subpath = append(subpath, k)
		}

		log.WithFields(log.Fields{
			"path":    fullpath.String(),
			"subpath": subpath,
		}).Debug("Install path handler")

		// path handler supports GET method only
		// a path handler displays only the group of available sub-path
		// hence we include the slash at the end
		r.Methods("GET").Path(fullpath.String()).Handler(
			&PathHandler{SubPath: subpath})
	}

	// go deep
	for k, n := range t.Nodes {
		path = append(path, k)
		traverseDepth(r, n, path)
		path = path[0 : len(path)-1]
	}
}
