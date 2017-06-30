package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// RegexConverter runs a regex on a single line
// with multi-line, the same regex is applied to each line
// example: /proc/version contains 5 parts:
// 1. os type
// 2. kernel release
// 3. user name
// 4. gcc version
// 5. kernel version
type RegexConverter struct {
	name      string
	regex     *regexp.Regexp
	title     []string // use the supplied title
	multiline bool
}

// NewRegexConverter
func NewRegexConverter(n string, x string, title []string, ml bool) *RegexConverter {
	return &RegexConverter{
		name:      n,
		regex:     regexp.MustCompile(x), // panic if failed
		title:     title,
		multiline: ml,
	}
}

// Name
func (c *RegexConverter) Name() string {
	return c.name
}

// ConvertStream
func (c *RegexConverter) ConvertStream(r io.Reader) (interface{}, error) {
	line := 0
	scanner := bufio.NewScanner(r)

	data := []map[string]string{}

	for scanner.Scan() {
		l := scanner.Text()

		dataLine := map[string]string{}

		groups := c.regex.FindStringSubmatch(l)

		log.WithFields(log.Fields{
			"title":  len(c.title),
			"groups": len(groups),
		}).Debug("Convert regex line")

		for i, g := range groups {
			log.WithFields(log.Fields{
				"idx": i,
			}).Debug(g)

			val := strings.Trim(g, " \t\n")

			if i == 0 {
				dataLine["full"] = val
			} else {
				if len(c.title) == 0 {
					dataLine[strconv.Itoa(i)] = val
				} else {
					dataLine[c.title[i-1]] = val
				}
			}
		}

		if line == 0 && !c.multiline {
			return dataLine, nil
		}

		data = append(data, dataLine)
		line++
	}

	return data, nil
}
