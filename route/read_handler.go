package route

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/source"
	"github.com/howardplus/lirest/util"
	"net/http"
)

type ReadHandler struct {
	Input  describe.DescriptionInput
	Output []describe.DescriptionOutputDesc
}

func (h *ReadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
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
	var conv source.Converter
	if format.Type == "separator" {
		conv = source.NewSeparatorConverter(format.Delimiter, format.Multiline, format.Multisection)
	}

	// read data from source
	var extractor source.Extractor
	if inputSource.Type == "filesystem" {
		extractor = source.NewFileSystemExtractor(inputSource.Path)
	} else {
		encoder.Encode(util.NamedError{Str: "Internal error: unknown input type"})
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
