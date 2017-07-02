package source

import (
	"github.com/howardplus/lirest/describe"
)

type Validator interface {
	// Validate data format
	Validate(data string) error
}

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
