package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"os"
)

const (
	ModeReadOnly  = iota // read-only value
	ModeReadWrite        // writable value
)

// AccessMode: either read or write
type AccessMode int

// ProcFSExtractor
type ProcFSExtractor struct {
	path string
}

// NewProcFSExtractor
func NewProcFSExtractor(path string) *ProcFSExtractor {
	return &ProcFSExtractor{path: path}
}

// Extract implements the Extractor interface
func (e *ProcFSExtractor) Extract(conv Converter) (interface{}, error) {
	// open file from path
	f, err := os.Open(e.path)
	if err != nil {
		return nil, &util.NamedError{Str: "Failed to open system path"}
	}
	defer f.Close()

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Extract from file system")

	// give it to the converter
	result, err := conv.ConvertStream(f)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Convert successful")

	return result, nil
}
