package source

import (
	"bufio"
	"bytes"
	log "github.com/Sirupsen/logrus"
	"io"
)

// AsisConverter
// "As-is" converter takes all the data as is
// this is the default converter
type AsisConverter struct {
	name      string
	multiline bool
}

// NewAsisConverter
func NewAsisConverter(n string, ml bool) *AsisConverter {
	return &AsisConverter{name: n, multiline: ml}
}

// Name
func (c *AsisConverter) Name() string {
	return c.name
}

// ConvertStream
func (c *AsisConverter) ConvertStream(r io.Reader) (interface{}, error) {
	line := 0
	scanner := bufio.NewScanner(r)

	var data bytes.Buffer

	for scanner.Scan() {
		l := scanner.Text()
		data.WriteString(l)
		line++
	}

	log.WithFields(log.Fields{
		"line": line,
	}).Debug("Convert as-is")

	return data.String(), nil
}
