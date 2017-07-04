package inject

import (
	"github.com/howardplus/lirest/util"
	"strconv"
	"strings"
)

// IntValidator validates integer based on min and max value
type IntValidator struct {
	min       int64
	max       int64
	multiline bool
}

// NewIntValidator creates a new integer validator
func NewIntValidator(min int64, max int64, multiline bool) *IntValidator {
	return &IntValidator{
		min:       min,
		max:       max,
		multiline: multiline,
	}
}

// Validate that checks if the value falls in the specified range
func (v *IntValidator) Validate(data string) error {

	var lines []string
	if v.multiline {
		lines = strings.Split(data, "\n")
	} else {
		lines = make([]string, 1, 1)
		lines[0] = data
	}

	for _, l := range lines {
		if i, err := strconv.ParseInt(l, 10, 64); err != nil {
			return util.NewError("NaN")
		} else if i > v.max || i < v.min {
			return util.NewError("Value out of range")
		}
	}

	return nil
}
