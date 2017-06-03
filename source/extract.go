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
	Extract(vars map[string]string) (map[string]interface{}, error)
}

// NewExtractor create a new extractor based on the description
func NewExtractor(s describe.DescriptionSource, c Converter) (Extractor, error) {
	var extractor Extractor

	switch s.Type {
	case "procfs", "sysfs", "sysctl":
		extractor = NewGenericExtractor(s.Path, c)
	case "command":
		extractor = NewCommandExtractor(s.Command, c)
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

func (e *GenericExtractor) Extract(vars map[string]string) (map[string]interface{}, error) {
	log.WithFields(log.Fields{
		"path": e.path,
		"vars": vars,
	}).Debug("Extract from file system")

	// create path from variables
	path, err := util.FillVars(e.path, vars)
	if err != nil {
		return nil, util.NewError("Failed to generate path")
	}

	// open file from path
	f, err := os.Open(path)
	if err != nil {
		return nil, util.NewError("Failed to open system path")
	}
	defer f.Close()

	// TODO: verify the rw format on this path

	// give it to the converter
	data, err := e.conv.ConvertStream(f)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Convert successful")

	return map[string]interface{}{
		"name": e.conv.Name(),
		"data": data,
	}, nil
}
