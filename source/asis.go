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

// ConvertLine
func (c *AsisConverter) ConvertLine(in string) (string, interface{}, error) {
	return "", in, nil
}

// ConvertStream
func (c *AsisConverter) ConvertStream(r io.Reader) (map[string]interface{}, error) {
	line := 0
	scanner := bufio.NewScanner(r)

	var data bytes.Buffer

	for scanner.Scan() {
		l := scanner.Text()
		_, v, _ := c.ConvertLine(l)
		data.WriteString(v.(string))
		line++
	}

	log.WithFields(log.Fields{
		"line": line,
	}).Debug("Convert as-is")

	return map[string]interface{}{
		"name": c.name,
		"data": data.String(),
	}, nil
}
