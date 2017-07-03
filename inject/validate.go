package inject

import (
	"github.com/howardplus/lirest/describe"
)

type Validator interface {
	// Validate data format
	Validate(data string) error
}

// NewValidator create a validator to validate data for injection
func NewValidator(format describe.DescriptionWriteFormat) Validator {
	switch format.Type {
	case "regex":
		return NewRegexValidator(format.Regex, format.Multiline)
	case "int":
		return NewIntValidator(format.Min, format.Max, format.Multiline)
	case "asis":
		return NewAsisValidator()
	}

	// by default, fail all validation
	return NewFailValidator()
}
