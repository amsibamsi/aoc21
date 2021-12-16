package d15p2

import (
	"math"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

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
	search := [][2]int{{0, 0}}
	for len(search) > 0 {
		p := search[0]
		search = search[1:]
		for _, d := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			next := [2]int{p[0] + d[0], p[1] + d[1]}
			risk := risks[p] + levels[next]
			if next == p {
				continue
			}
			if _, ok := levels[next]; !ok {
				continue
			}
			if risk >= max {
				continue
			}
			if r, ok := risks[next]; ok && risk >= r {
				continue
			}
			risks[next] = risk
			if next == end {
				if risk < max {
					max = risk
				}
				continue
			}
			search = append(search, next)
		}
	}
	return strconv.Itoa(risks[end]), nil
}
