package util

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

// FillVars fill variable values into s
func FillVars(s string, vars map[string]string) (string, error) {
	for k, v := range vars {
		start := strings.Index(s, "{"+k+"}")
		end := start + 2 + len(k)

		// the command or path may not use the variable
		if start != -1 {
			s = s[:start] + v + s[end:]
		}
	}

	return s, nil
}

// PathAddType add the type string into the path string
// this is used to display path info
func PathAddType(path string, vars map[string]string) string {
	last := 0
	for k, v := range vars {
		if start := strings.Index(path[last:], "{"+k+"}"); start != -1 {
			end := start + 2 + len(k)
			path = path[:start+last] + "{" + k + ":" + v + "}" + path[end+last:]

			last += start + 1

			log.WithFields(log.Fields{
				"key":  k,
				"val":  v,
				"path": path,
				"last": last,
			}).Debug("Add var to path")

		} else {
			break
		}
	}
	return path
}
