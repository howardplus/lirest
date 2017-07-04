package inject

import (
	"github.com/howardplus/lirest/describe"
)

// Validate data format before injected
type Validator interface {
	Validate(data string) error
}

// NewValidator create a new validator based on type
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
