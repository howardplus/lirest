package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"os"
)

const (
	ModeReadOnly = iota
	ModeReadWrite
)

type AccessMode int

type FileSystemExtractor struct {
	path string
}

func NewFileSystemExtractor(path string) *FileSystemExtractor {
	return &FileSystemExtractor{path: path}
}

// implements the Extractor interface
func (e *FileSystemExtractor) Extract(conv Converter) (interface{}, error) {
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
