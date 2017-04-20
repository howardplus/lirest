package util

type NamedError struct {
	Str string `json:"error"`
}

// Error
// use the string in the error, also support json format
func (e *NamedError) Error() string {
	return e.Str
}
