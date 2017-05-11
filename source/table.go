package source

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"io"
	"strconv"
	"strings"
)

const colTitleOther = "Other"

// TableConverter converts a table-like output
// example: /proc/interrupts
type TableConverter struct {
	name       string
	hasTitle   bool // title goes to the top (row 0)
	hasHeading bool // heading goes to the left (col 0)
}

// NewTableConverter create a new table separator
func NewTableConverter(n string, t bool, h bool) *TableConverter {
	return &TableConverter{name: n, hasTitle: t, hasHeading: h}
}

// ConvertLine converts a line into a row
func (c *TableConverter) ConvertLine(in string) (key string, value interface{}, err error) {

	log.WithFields(log.Fields{
		"line": in,
	}).Debug("Table convert line")

	fields := strings.Fields(in)
	if len(fields) < 1 {
		return "", nil, util.NewError("Insufficient fields in a line")
	}

	// just return the fields as columns
	return "", fields, nil
}

// ConvertStream streams input, read line by line
func (c *TableConverter) ConvertStream(r io.Reader) (map[string]interface{}, error) {

	output := make(map[string]interface{}, 0)

	line := 0
	scanner := bufio.NewScanner(r)

	rows := make([]map[string]string, 0)

	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \t")

		if c.hasTitle && line == 0 {
			// get title line
			line++
			output["title"] = strings.Fields(l)
			continue
		}

		row := make(map[string]string, 0)
		heading := ""
		if c.hasHeading {
			// with heading, there needs to be at least 2 fields
			if fields := strings.Fields(l); len(fields) >= 2 {
				heading = fields[0]

				log.WithFields(log.Fields{
					"heading": heading,
				}).Debug("With heading")

			} else {
				return nil, util.NewError("Heading expected on line")
			}
		}

		row["heading"] = heading

		// strip off the heading and send it to convertline
		tmp := strings.Trim(strings.TrimPrefix(l, heading), " \t")

		// convert line by line
		_, v, err := c.ConvertLine(tmp)
		if err != nil {
			return nil, err
		}

		// walk columns
		for i, elem := range v.([]string) {
			if title := output["title"].([]string); title != nil {
				// retrieve cols that we know
				if i < len(title) {
					row[title[i]] = elem
				} else {
					if _, found := row[colTitleOther]; !found {
						row[colTitleOther] = elem
					} else {
						row[colTitleOther] = row[colTitleOther] + " " + elem
					}
				}
			} else {
				// no title, each field is a column
				row[strconv.Itoa(i)] = elem
			}
		}

		// add to rows
		rows = append(rows, row)
		line++
	}

	output["rows"] = rows

	return map[string]interface{}{
		"name": c.name,
		"data": output,
	}, nil
}
