package d15p1

import (
	"math"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	levels := map[[2]int]int{}
	lines := aoc21.ToLines(input)
	for y, line := range lines {
		for x, level := range line {
			num, err := strconv.Atoi(string(level))
			if err != nil {
				return "", err
			}
			levels[[2]int{x, y}] = num
		}
	}
	end := [2]int{len(lines[0]) - 1, len(lines) - 1}
	risks := map[[2]int]int{}
	max := math.MaxInt
	search([2]int{0, 0}, end, levels, risks, 0, &max)
	return strconv.Itoa(risks[end]), nil
}

func search(p, end [2]int, levels map[[2]int]int, risks map[[2]int]int, risk int, max *int) {
	if _, ok := levels[p]; !ok {
		return
	}
	if risk >= *max {
		return
	}
	if r, ok := risks[p]; ok && risk >= r {
		return
	}
	risks[p] = risk
	if p == end {
		if risk < *max {
			*max = risk
		}
		return
	}
	for _, d := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := [2]int{p[0] + d[0], p[1] + d[1]}
		search(next, end, levels, risks, risk+levels[next], max)
	}
}
