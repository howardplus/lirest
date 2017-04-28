package util

type Result struct {
	Str string `json:"result"`
}

func NewResultOk() *Result {
	return &Result{Str: "ok"}
}
