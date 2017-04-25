package source

import (
	_ "github.com/howardplus/lirest/describe"
)

// Formatter returns a string that is
// piped to the source
type Formatter interface {
	Format(conv Converter) (string, error)
}

// NewFormatter creates a new formatter based on the description
