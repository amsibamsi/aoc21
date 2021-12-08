package d08p1

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	var count [8]int
	for _, line := range aoc21.ToLines(input) {
		output := strings.Split(line, " | ")[1]
		for _, digit := range strings.Split(output, " ") {
			count[len(digit)]++
		}
	}
	return strconv.Itoa(count[2] + count[3] + count[4] + count[7]), nil
}
