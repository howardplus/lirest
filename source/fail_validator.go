package source

import (
	"github.com/howardplus/lirest/util"
)

// FailValidator always return fail
type FailValidator struct {
}

func NewFailValidator() *FailValidator {
	return &FailValidator{}
}

// Validate that return failure
func (v *FailValidator) Validate(data string) error {
	return util.NewError("Validation error")
}
