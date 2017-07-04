package util

// Result is a structure to return result in json format
type Result struct {
	Str string `json:"result"`
}

// NewResultOk returns ok
func NewResultOk() *Result {
	return &Result{Str: "ok"}
}
