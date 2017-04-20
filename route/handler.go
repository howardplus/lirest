package route

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// SimpleHandler
// a placeholder handler that only logs the path and method
func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
	}).Debug("SimpleHandler")
}
