package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/util"
	"net/http"
	"strconv"
)

type GetHandler func(map[string]string) interface{}
type SetHandler func(map[string]string, interface{}) error
type DelHandler func(map[string]string) error
type NewHandler func(map[string]string) error

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

func GenerateCommandRoutes(r *mux.Router) error {

	s := r.PathPrefix("/cmd").Subrouter()

	// create command routes
	cmds := []string{}
	for path, v := range CommandRoutes {
		s.Methods("GET", "PUT", "POST", "DELETE").Path(path).Handler(&v)
		cmds = append(cmds, path)
	}

	// create top level command list
	s.Methods("GET").Path("/").Handler(&CommandRootHandler{
		cmds: cmds,
	})

	return nil
}

type CommandRootHandler struct {
	cmds []string
}

func (h *CommandRootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	if config.GetConfig().Pretty {
		encoder.SetIndent("", "  ")
	}

	encoder.Encode(map[string]interface{}{
		"name": "commands",
		"data": h.cmds,
	})
}

type CommandHandler struct {
	name string
	get  GetHandler
	set  SetHandler
	del  DelHandler
	new  NewHandler
}

func (h *CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"method": r.Method,
		"name":   h.name,
	}).Info("Run command")

	encoder := json.NewEncoder(w)

	if config.GetConfig().Pretty {
		encoder.SetIndent("", "  ")
	}

	switch r.Method {
	case "GET":
		if h.get == nil {
			encoder.Encode(util.NewError("Unsupported"))
		}

		encoder.Encode(map[string]interface{}{
			"name": h.name,
			"data": h.get(mux.Vars(r)),
		})

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

	case "POST":
		if h.new == nil {
			encoder.Encode(util.NewError("Unsupported"))
			return
		}

	case "DELETE":
		if h.del == nil {
			encoder.Encode(util.NewError("Unsupported"))
			return
		}
	}

}
