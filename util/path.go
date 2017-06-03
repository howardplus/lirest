package util

import (
	"strings"
)

func PathFillVars(path string, vars map[string]string) (string, error) {
	for k, v := range vars {
		// TODO: add more error checking and utest
		start := strings.Index(path, "{"+k+"}")
		end := start + 2 + len(k)
		path = path[:start] + v + path[end:]
	}

	return path, nil
}

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
