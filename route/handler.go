package route

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
	}).Debug("SimpleHandler")
}
