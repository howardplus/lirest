package util

import (
	"strings"
)

// FillVars fill variable values into s
func FillVars(s string, vars map[string]string) (string, error) {
	for k, v := range vars {
		// TODO: add more error checking and utest
		start := strings.Index(s, "{"+k+"}")
		end := start + 2 + len(k)
		s = s[:start] + v + s[end:]
	}

	return s, nil
}

// PathAddType add the type string into the path string
// this is used to display path info
func PathAddType(path string, vars map[string]string) string {
	last := 0
	for k, v := range vars {
		if start := strings.Index(path[last:], "{"+k+"}"); start != -1 {
			last += start + 1
			end := start + 2 + len(k)
			path = path[:start] + "{" + k + ":" + v + "}" + path[end:]
		} else {
			break
		}
	}
	return path
}
