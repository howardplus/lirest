package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"io"
	"strings"
)

// SeparatorConverter:
// separated by a character
// typical example is ':' separated
// key: value
type SeparatorConverter struct {
	name         string
	sep          string
	multiline    bool
	multisection bool
}

// NewSeparatorConverter
// create a new separator
func NewSeparatorConverter(n string, s string, ml bool, ms bool) *SeparatorConverter {
	return &SeparatorConverter{name: n, sep: s, multiline: ml, multisection: ms}
}

// ConvertLine
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

// ConvertStream
// stream input, read line by line
func (c *SeparatorConverter) ConvertStream(r io.Reader) (map[string]interface{}, error) {

	output := make([]map[string]interface{}, 0)

	section := 0
	line := 0

	// per-section output
	outputSection := make(map[string]interface{}, 10)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()

		if l == "" {
			log.Debug("empty line")
			// for multi-section, an empty marks the end of a section
			output = append(output, outputSection)
			// recreate the map
			outputSection = make(map[string]interface{}, 10)
			section++
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
		outputSection[k] = v

		if !c.multiline && line == 0 {
			// single line source, done
			return map[string]interface{}{
				"name": c.name,
				"data": outputSection,
			}, nil
		}

		line++
	}

	// there may not be an empty line at all
	// in this case, take the entire section as output
	if len(output) == 0 {
		return map[string]interface{}{
			"name": c.name,
			"data": outputSection,
		}, nil
	}

	// otherwise, prepare the multi-section output
	return map[string]interface{}{
		"name": c.name,
		"data": output,
	}, nil
}
