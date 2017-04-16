package route

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateSysRoute(parent *mux.Router) http.Handler {
	log.Info("Create system route")

	r := parent.PathPrefix("/sys").Subrouter()
	r.Methods("GET").Path("/").HandlerFunc(sysHandler)

	return r
}

func sysHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(TagHeaderName) == TagInfo {
		log.Info("sys info handler")
	} else if r.Header.Get(TagHeaderName) == TagMan {
		log.Info("sys man handler")
	} else {
		log.Info("sys handler")
	}
}
