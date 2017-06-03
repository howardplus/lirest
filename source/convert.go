package source

import (
	"github.com/howardplus/lirest/describe"
	"io"
	"strings"
)

// Converter
// generic convert interface
type Converter interface {

	// Name returns the name for the converter
	Name() string

	// stream conversion takes an io.Reader and convert
	// line by line into a map
	ConvertStream(r io.Reader) (interface{}, error)
}

// ConvertValue
// convert a value into proper formats
// TODO: just a plain string for now
func ConvertValue(s string) interface{} {
	return strings.Trim(s, " \t")
}

// NewConverter
// Create a converter based on its type
func NewConverter(name string, format describe.DescriptionFormat) Converter {
	switch format.Type {
	case "separator":
		return NewSeparatorConverter(name, format.Delimiter, format.Multiline, format.Multisection)
	case "list":
		return NewListConverter(name, format.Header, format.Title, format.Multiline)
	case "regex":
		return NewRegexConverter(name, format.Regex, format.Title, format.Multiline)
	case "table":
		return NewTableConverter(name, format.HasTitle, format.HasHeading)
	case "asis":
		return NewAsisConverter(name, format.Multiline)
	case "line-by-line":
		return NewLBLConverter(name, format.Multisection)
	}

	// default is as-is
	return NewAsisConverter(name, format.Multiline)
}
