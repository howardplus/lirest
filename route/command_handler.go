package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/util"
	"net/http"
	"strconv"
)

type GetHandler func(map[string]string) interface{}
type SetHandler func(map[string]string, interface{}) error

var CommandRoutes map[string]CommandHandler = map[string]CommandHandler{
	"/jobs": CommandHandler{
		name: "jobs",
		get: func(vars map[string]string) interface{} {
			return inject.RequestJobs(0)
		},
	},
	"/jobs/{n:[0-9]+}": CommandHandler{
		name: "jobs",
		get: func(vars map[string]string) interface{} {
			n, _ := strconv.Atoi(vars["n"])
			return inject.RequestJobs(n)
		},
	},
}

type CommandHandler struct {
	name string
	get  GetHandler
	set  SetHandler
}

func GenerateCommandRoutes(r *mux.Router) error {

	s := r.PathPrefix("/cmd").Subrouter()

	for path, v := range CommandRoutes {
		s.Methods("GET", "PUT").Path(path).Handler(&CommandHandler{
			name: v.name,
			get:  v.get,
			set:  v.set,
		})
	}

	return nil
}

func (h *CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"method": r.Method,
		"name":   h.name,
	}).Info("Run command")

	encoder := json.NewEncoder(w)
	switch r.Method {
	case "GET":
		if h.get == nil {
			encoder.Encode(util.NewError("Unsupported"))
		}

		// add title
		data := make(map[string]interface{}, 1)
		data["name"] = h.name
		data["data"] = h.get(mux.Vars(r))

		// encode it
		encoder.SetIndent("", "  ")
		encoder.Encode(data)

	case "PUT":
		if h.set == nil {
			encoder.Encode(util.NewError("Unsupported"))
			return
		}

		// TODO: how to pass in data
		if err := h.set(mux.Vars(r), nil); err != nil {
			encoder.Encode(err)
		} else {
			encoder.Encode(util.NewResultOk())
		}
	}
}
