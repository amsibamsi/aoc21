package d12p2

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	conns := make(map[string][]string)
	for _, line := range aoc21.ToLines(input) {
		c := strings.Split(line, "-")
		mapAppend(conns, c[0], c[1])
		mapAppend(conns, c[1], c[0])
	}
	paths := search([]string{"start"}, true, conns)
	return strconv.Itoa(len(paths)), nil
}

func search(path []string, canRepeatSmall bool, conns map[string][]string) [][]string {
	loc := path[len(path)-1]
	if loc == "end" {
		newPath := make([]string, len(path))
		copy(newPath, path)
		return [][]string{newPath}
	}
	paths := make([][]string, 0)
	for _, c := range conns[loc] {
		isRepeatedSmall := !isUpper(c) && c != "start" && contains(path, c)
		if !contains(path, c) || isUpper(c) || canRepeatSmall && isRepeatedSmall {
			newPaths := search(append(path, c), canRepeatSmall && !isRepeatedSmall, conns)
			paths = append(paths, newPaths...)
		}
	}
	return paths
}

func isUpper(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) || !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}

func contains(path []string, cave string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

func mapAppend(m map[string][]string, k, v string) {
	_, ok := m[k]
	if ok {
		m[k] = append(m[k], v)
	} else {
		m[k] = []string{v}
	}
}
