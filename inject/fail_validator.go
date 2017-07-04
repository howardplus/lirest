package inject

import (
	"github.com/howardplus/lirest/util"
)

// FailValidator always return fail
type FailValidator struct {
}

// NewFailValidator creates a new validator that always fails
func NewFailValidator() *FailValidator {
	return &FailValidator{}
}

// Validate that return failure
func (v *FailValidator) Validate(data string) error {
	return util.NewError("Validation error")
}
