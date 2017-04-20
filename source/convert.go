package source

import (
	"github.com/howardplus/lirest/describe"
	"io"
	"strings"
)

// Converter
// generic convert interface
type Converter interface {
	// single line converter takes an input line
	// and convert into key-value pair
	ConvertLine(in string) (key string, value interface{}, err error)

	// stream conversion takes an io.Reader and convert
	// line by line into a map
	ConvertStream(r io.Reader) (result map[string]interface{}, err error)
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
	case "asis":
		return NewAsisConverter(name, format.Multiline)
	}

	// default is as-is
	return NewAsisConverter(name, format.Multiline)
}
