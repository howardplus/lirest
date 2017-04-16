package source

import (
	"io"
	"strings"
)

// generic convert interface
type Converter interface {
	// single line converter takes an input line
	// and convert into key-value pair
	ConvertLine(in string) (key string, value interface{}, err error)

	// stream conversion takes an io.Reader and convert
	// line by line into a map
	ConvertStream(r io.Reader) (result map[string]interface{}, err error)
}

// Convert Value into proper formats
// TODO: just a plain string for now
func ConvertValue(s string) interface{} {
	return strings.Trim(s, " \t")
}
