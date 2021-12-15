package d15p2

import (
	"math"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

// TODO: Terribly slow, ~400s on single core. But hey, memory usage is
// awesome!
func Solve(input string) (string, error) {
	levels := map[[2]int]int{}
	lines := aoc21.ToLines(input)
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, level := range line {
			num, err := strconv.Atoi(string(level))
			if err != nil {
				return "", err
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					l := (num+i+j-1)%9 + 1
					levels[[2]int{x + i*width, y + j*height}] = l
				}
			}
		}
	}
	end := [2]int{5*width - 1, 5*height - 1}
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
