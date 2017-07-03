package inject

// AsisValidator always return ok
type AsisValidator struct {
}

func NewAsisValidator() *AsisValidator {
	return &AsisValidator{}
}

// Validate that return ok
func (v *AsisValidator) Validate(data string) error {
	return nil
}
