package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// regex converter runs a regex on a single line
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

func NewRegexConverter(n string, x string, title []string, ml bool) *RegexConverter {
	return &RegexConverter{
		name:      n,
		regex:     regexp.MustCompile(x), // panic if failed
		title:     title,
		multiline: ml,
	}
}

func (c *RegexConverter) ConvertLine(in string) (string, interface{}, error) {
	// just let regex do its thing
	groups := c.regex.FindStringSubmatch(in)
	return "", groups, nil
}

func (c *RegexConverter) ConvertStream(r io.Reader) (map[string]interface{}, error) {
	line := 0
	scanner := bufio.NewScanner(r)

	data := []map[string]string{}

	for scanner.Scan() {
		l := scanner.Text()

		data_l := map[string]string{}

		_, v, _ := c.ConvertLine(l)

		log.WithFields(log.Fields{
			"title":  len(c.title),
			"groups": len(v.([]string)),
		}).Debug("Convert regex line")

		for i, g := range v.([]string) {
			log.WithFields(log.Fields{
				"idx": i,
			}).Debug(g)

			val := strings.Trim(g, " \t\n")

			if i == 0 {
				data_l["full"] = val
			} else {
				if len(c.title) == 0 {
					data_l[strconv.Itoa(i)] = val
				} else {
					data_l[c.title[i-1]] = val
				}
			}
		}

		if line == 0 && !c.multiline {
			return map[string]interface{}{
				"name": c.name,
				"data": data_l,
			}, nil
		}

		data = append(data, data_l)
		line++
	}

	return map[string]interface{}{
		"name": c.name,
		"data": data,
	}, nil
}