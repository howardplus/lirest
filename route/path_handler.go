package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/config"
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

	encoder := json.NewEncoder(w)
	if config.GetConfig().Pretty {
		encoder.SetIndent("", "  ")
	}

	// encode it
	encoder.Encode(map[string]interface{}{
		"subpath": subpath,
	})
}
