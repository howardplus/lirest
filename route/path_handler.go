package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type PathHandler struct {
	SubPath []string
}

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
