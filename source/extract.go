package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
	"os"
)

// Extractor returns a generic data based
// on the converter.
// An object that implements the Extract() interface needs
// to know where to get the data, which then feeds to the
// converter.
type Extractor interface {
	Extract() (interface{}, error)
}

// NewExtractor create a new extractor based on the description
func NewExtractor(s describe.DescriptionSource, c Converter) (Extractor, error) {
	var extractor Extractor

	switch s.Type {
	case "procfs", "sysfs", "sysctl":
		extractor = NewGenericExtractor(s.Path, c)
	}

	// found an extractor, use it
	if extractor != nil {
		return extractor, nil
	}

	// return error on default
	return nil, util.NewError("Internal error: unknown input type")
}

// GenericExtractor
type GenericExtractor struct {
	path string
	conv Converter
}

// GenericExtractor extract data from reading from a file
// use this until it's not enough
func NewGenericExtractor(path string, conv Converter) *GenericExtractor {
	return &GenericExtractor{path: path, conv: conv}
}

func (e *GenericExtractor) Extract() (interface{}, error) {
	// open file from path
	f, err := os.Open(e.path)
	if err != nil {
		return nil, util.NewError("Failed to open system path")
	}
	defer f.Close()

	// TODO: verify the rw format on this path

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Extract from file system")

	// give it to the converter
	result, err := e.conv.ConvertStream(f)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Convert successful")

	return result, nil
}
