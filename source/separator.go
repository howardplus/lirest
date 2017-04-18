package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"io"
	"strings"
)

// separated by a character
// typical example is ':' separated
// key: value
type SeparatorConverter struct {
	sep          string
	multiline    bool
	multisection bool
}

// create a new separator
func NewSeparatorConverter(s string, ml bool, ms bool) *SeparatorConverter {
	return &SeparatorConverter{sep: s, multiline: ml, multisection: ms}
}

// single line, seperated by the separator
func (c *SeparatorConverter) ConvertLine(in string) (key string, value interface{}, err error) {
	parts := strings.Split(strings.Trim(in, " \t"), c.sep)

	if len(parts) != 2 {
		return "", nil, &util.NamedError{Str: "Insufficient parts of a separator line"}
	}

	// now gets the key and value
	key = strings.Trim(parts[0], " \t")

	// the value can be in many forms
	value = ConvertValue(parts[1])

	return key, value, nil
}

// stream input, read line by line
func (c *SeparatorConverter) ConvertStream(r io.Reader) (map[string]interface{}, error) {

	output := make(map[string]interface{}, 10)
	line := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// TODO: multi-section requires a top level section title
		/*
			if c.multisection {
				return nil, nil
			}
		*/

		l := scanner.Text()

		if l == "" {
			continue
		}

		// convert line by line
		k, v, err := c.ConvertLine(l)
		if err != nil {
			return nil, err
		}

		log.WithFields(log.Fields{
			"key":   k,
			"value": v,
		}).Debug("converted")

		// retrieve the data
		line++
		output[k] = v

		if !c.multiline && line == 1 {
			// single line source, done
			return output, nil
		}
	}

	// prepare the output
	return output, nil
}
