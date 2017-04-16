package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/output"
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
	subpath := []output.SubPath{}
	for _, sp := range h.SubPath {
		subpath = append(subpath, output.SubPath{Path: sp})
	}

	// encode it
	json.NewEncoder(w).Encode(subpath)
}
