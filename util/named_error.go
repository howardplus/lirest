package util

// NamedError is similar to "errors" package except
// that it can be encoded as json
type NamedError struct {
	Str string `json:"error"`
}

// NewError creates an error with specified string
func NewError(s string) error {
	return &NamedError{Str: s}
}

// Error use the string in the error, also support json format
func (e *NamedError) Error() string {
	return e.Str
}
