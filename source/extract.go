package source

import (
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
)

// Extractor returns a generic data based
// on the converter.
// An object that implements the Extract() interface needs
// to know where to get the data, which then feeds to the
// converter.
type Extractor interface {
	Extract(conv Converter) (interface{}, error)
}

// Create a new extractor based on the description
func NewExtractor(source describe.DescriptionSource) (Extractor, error) {
	var extractor Extractor

	switch source.Type {
	case "procfs":
		extractor = NewProcFSExtractor(source.Path)
	}

	// found an extractor, use it
	if extractor != nil {
		return extractor, nil
	}

	// return error on default
	return nil, &util.NamedError{Str: "Internal error: unknown input type"}
}
