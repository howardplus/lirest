package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/util"
	"net/http"
)

type GetHandler func() interface{}
type SetHandler func(interface{}) error

var CommandRoutes map[string]CommandHandler = map[string]CommandHandler{
	"/jobs": CommandHandler{
		get: func() interface{} {
			return inject.RequestJobs(0)
		},
	},
}

type CommandHandler struct {
	cmd string
	get GetHandler
	set SetHandler
}

func GenerateCommandRoutes(r *mux.Router) error {

	s := r.PathPrefix("/cmd").Subrouter()

	for k, v := range CommandRoutes {
		s.Methods("GET", "PUT").Path(k).Handler(&CommandHandler{
			cmd: k,
			get: v.get,
			set: v.set,
		})
	}

	return nil
}

func (h *CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"method": r.Method,
		"cmd":    h.cmd,
	}).Info("Run command")

	encoder := json.NewEncoder(w)
	switch r.Method {
	case "GET":
		if h.get == nil {
			encoder.Encode(util.NewError("Unsupported"))
		}

		// add title
		data := make(map[string]interface{}, 1)
		data["name"] = h.cmd
		data["data"] = h.get()

		// encode it
		encoder.SetIndent("", "  ")
		encoder.Encode(data)

	case "PUT":
		if h.set == nil {
			encoder.Encode(util.NewError("Unsupported"))
			return
		}

		if err := h.set(nil); err != nil {
			encoder.Encode(err)
		} else {
			encoder.Encode(util.NewResultOk())
		}
	}
}
