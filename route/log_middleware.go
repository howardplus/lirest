package route

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"time"
)

func LogMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	now := time.Now()
	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"tag":    r.Header.Get(TagHeaderName),
	}).Info("Access: ", now.Year(), "-", now.Month(), "-", now.Day(), " ", now.Hour(), ":", now.Minute(), ":", now.Second())

	next(w, r)
}
