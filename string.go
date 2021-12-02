package aoc21

import "strings"

func ToLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}
