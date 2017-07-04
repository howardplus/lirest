package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"io"
	"strings"
)

// LBLConverter converts data line by line
// "line-by-line" converter spits out data one per line
type LBLConverter struct {
	name         string
	multisection bool
}

// NewLBLConverter creates a new line-by-line converter
func NewLBLConverter(n string, ms bool) *LBLConverter {
	return &LBLConverter{name: n, multisection: ms}
}

// Name returns the name of the converter
func (c *LBLConverter) Name() string {
	return c.name
}

// ConvertStream converts data line-by-line
func (c *LBLConverter) ConvertStream(r io.Reader) (interface{}, error) {
	line := 0
	scanner := bufio.NewScanner(r)

	var data []string = make([]string, 0, 0)

	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \t")
		data = append(data, l)
		line++
	}

	log.WithFields(log.Fields{
		"line": line,
	}).Debug("Convert line by line")

	return data, nil
}
