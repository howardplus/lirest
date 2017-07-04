package inject

// AsisValidator always return ok
type AsisValidator struct {
}

// NewAsisValidator creates a new as-is validator
func NewAsisValidator() *AsisValidator {
	return &AsisValidator{}
}

// Validate that always return ok
func (v *AsisValidator) Validate(data string) error {
	return nil
}
