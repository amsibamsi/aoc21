package d09p1

import (
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	heights := [][]int{}
	for _, textRow := range aoc21.ToLines(input) {
		row := []int{}
		for _, num := range textRow {
			n, err := strconv.Atoi(string(num))
			if err != nil {
				return "", err
			}
			row = append(row, n)
		}
		heights = append(heights, row)
	}
	rlevels := 0
	for y := range heights {
		for x := range heights[y] {
			min := true
			for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				yn := y + d[0]
				xn := x + d[1]
				if yn >= 0 && yn < len(heights) && xn >= 0 && xn < len(heights[y]) {
					if heights[yn][xn] <= heights[y][x] {
						min = false
					}
				}
			}
			if min {
				rlevels += heights[y][x] + 1
			}
		}
	}
	return strconv.Itoa(rlevels), nil
}
