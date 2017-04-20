package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/source"
	"net/http"
)

type ReadHandler struct {
	Name   string
	Input  describe.DescriptionInput
	Output []describe.DescriptionOutputDesc
}

func (h *ReadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	inputSource := h.Input.Source
	format := h.Input.Format
	tag := r.Header.Get(TagHeaderName)

	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"type":   inputSource.Type,
		"tag":    tag,
	}).Debug("ReadHandler")

	// check if we are on tagged path
	if tag == TagInfo {
		// info tag describes the output format, which is how
		// the user uses the REST api
		encoder.Encode(h.Output)
		return
	} else if tag == TagMan {
		// man tag is free style using markdown language
		return
	}

	// create format converter
	conv := source.NewConverter(h.Name, format)

	// read data from source
	extractor, err := source.NewExtractor(inputSource)
	if err != nil {
		encoder.Encode(err)
		return
	}

	// extract the data
	output, err := extractor.Extract(conv)
	if err != nil {
		encoder.Encode(err)
		return
	}

	// all good, encode the data and send back
	encoder.Encode(output)
}
