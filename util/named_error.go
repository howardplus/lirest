package util

type NamedError struct {
	Str string `json:"error"`
}

func (e *NamedError) Error() string {
	return e.Str
}
