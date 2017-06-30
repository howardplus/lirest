package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"io"
	"strconv"
	"strings"
)

// ListConverter
// list are values that are listed in either
// a single line, or multi-line
// for example:
// /proc/uptime contains 2 values of a single line:
type ListConverter struct {
	name      string
	header    bool     // does this converter have header, which is the first line
	title     []string // use the supplied title
	multiline bool
}

// NewListConverter
// Create a list converter
func NewListConverter(n string, h bool, t []string, ml bool) *ListConverter {
	return &ListConverter{name: n, header: h, title: t, multiline: ml}
}

// Name
func (c *ListConverter) Name() string {
	return c.name
}

// ConvertStream
// Convert from a io.Reader
func (c *ListConverter) ConvertStream(r io.Reader) (interface{}, error) {

	// output is a slice of map of title to value
	output := []map[string]string{}
	title := []string{}

	line := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		outputLine := make(map[string]string, 10) // per line output

		l := scanner.Text()
		if l == "" {
			continue
		}

		log.WithFields(log.Fields{
			"line": line,
		}).Info("Add list")

		fields := strings.Fields(l)

		// use the title as key
		// example: /proc/swaps
		for col, val := range fields {

			// first line is stored as title
			if line == 0 && c.header {
				title = append(title, val)

				log.WithFields(log.Fields{
					"col": col,
					"val": val,
				}).Debug("Add title from header")
				continue
			}

			// no header, get title either from c.title
			// or just the col number
			if !c.header {
				if len(c.title) == 0 {
					title = append(title, strconv.Itoa(col))
				} else {
					title = append(title, c.title[col])
				}
			}

			log.WithFields(log.Fields{
				"col":   col,
				"title": title[col],
				"val":   val,
			}).Debug("Add Title")

			// output with title
			outputLine[title[col]] = val
		}

		if !c.multiline {
			return outputLine, nil
		} else if line != 0 || !c.header {
			output = append(output, outputLine)
		}

		// done. next line
		line++
	}

	return output, nil
}
