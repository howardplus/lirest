package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// PathHandler contains slice of next level subpath
type PathHandler struct {
	SubPath []string
}

// ServeHTTP
// the HTTP handler
func (h *PathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"method":  r.Method,
		"path":    r.URL.Path,
		"subpath": h.SubPath,
	}).Debug("PathHandler")

	// create response with links to subpath
	subpath := []string{}
	for _, sp := range h.SubPath {
		subpath = append(subpath, sp)
	}

	// add title
	data := make(map[string]interface{}, 1)
	data["name"] = "subpath"
	data["data"] = subpath

	// encode it
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}
