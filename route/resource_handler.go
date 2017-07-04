package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/howardplus/lirest/config"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/source"
	"github.com/howardplus/lirest/util"
	"net/http"
)

// ResourceHandler Describe a resource handler
type ResourceHandler struct {
	Name   string
	System describe.DescriptionSystem
	Api    describe.DescriptionApi
}

// input from POST/PUT
type userData struct {
	Data string `json:"data"`
}

// ServeHTTP
// ResourceHandler's HTTP handler function
func (h *ResourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := h.System.Source
	tag := r.Header.Get(TagHeaderName)
	vars := mux.Vars(r)

	encoder := json.NewEncoder(w)

	if config.GetConfig().Pretty {
		encoder.SetIndent("", "  ")
	}

	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"type":   s.Type,
		"tag":    tag,
		"vars":   vars,
	}).Debug("ResourceHandler serve")

	// check type compatibility of each var
	for _, v := range h.Api.Vars {
		varType := v.DataType
		if val, found := vars[v.Name]; found == false {
			encoder.Encode(util.NewError("Variable not found"))
			return
		} else if !describe.DescriptionVarValidate(val, varType) {
			encoder.Encode(util.NewError("Variable validation error"))
			return
		}
	}

	// handle based on method type
	switch r.Method {
	case "GET":

		// check if we are on tagged path
		if tag == TagInfo {
			// info tag describes the output format, which is how
			// the user uses the REST api
			encoder.Encode(map[string]interface{}{
				"path": h.Api.Path,
				"api":  h.Api.Descriptions,
			})
			return
		}

		// create format converter
		conv := source.NewConverter(h.Name, h.System.ReadFormat)

		// read data from source
		extractor, err := source.NewExtractor(s, conv, vars)
		if err != nil {
			encoder.Encode(err)
			return
		}

		// extract the data
		output, err := extractor.Extract()
		if err != nil {
			encoder.Encode(err)
			return
		}

		// all good, encode the data and send back
		encoder.Encode(map[string]interface{}{
			"name": output.Name,
			"time": output.Time,
			"data": output.Data,
		})

	case "PUT":
		// do not support tag
		if tag != "" {
			http.NotFound(w, r)
			return
		}

		// decode user input
		var data userData
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		if err := decoder.Decode(&data); err != nil {
			// TODO: need better http response
			http.NotFound(w, r)
			return
		}

		log.WithFields(log.Fields{
			"data": data,
		}).Info("Received user data")

		injector, err := inject.NewInjector(h.System.Source, h.System.WriteFormat)
		if err != nil {
			encoder.Encode(err)
			return
		}

		// perform the system setting
		if err := injector.Inject(data.Data); err != nil {
			encoder.Encode(err)
			return
		}

		// success
		encoder.Encode(util.NewResultOk())
	}
}
