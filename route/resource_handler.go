package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/source"
	"net/http"
)

// ResourceHandler Describe a resource handler
type ResourceHandler struct {
	Name   string
	System describe.DescriptionSystem
	Api    []describe.DescriptionApiDesc
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

	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"type":   s.Type,
		"tag":    tag,
	}).Debug("ResourceHandler serve")

	switch r.Method {
	case "GET":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		// check if we are on tagged path
		if tag == TagInfo {
			// info tag describes the output format, which is how
			// the user uses the REST api
			encoder.Encode(h.Api)
			return
		} else if tag == TagMan {
			// man tag is free style using markdown language
			// TODO
			return
		}

		// create format converter
		conv := source.NewConverter(h.Name, h.System.ReadFormat)

		// read data from source
		extractor, err := source.NewExtractor(s, conv)
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
		encoder.Encode(output)

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

		injector, err := source.NewInjector(h.System.Source, h.System.WriteFormat)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		// perform the system setting
		// TODO: sandbox
		_, err = injector.Inject(data.Data)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		// validate input format
		w.Write([]byte("200 OK\n"))
	}
}
