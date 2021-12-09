package d09p2

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
	lows := [][2]int{}
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
				lows = append(lows, [2]int{y, x})
			}
		}
	}
	top3 := [3]int{0, 0, 0}
	seen := make([][]bool, len(heights))
	for y := range heights {
		seen[y] = make([]bool, len(heights[y]))
		for x := range heights[y] {
			seen[y][x] = false
		}
	}
	for _, low := range lows {
		size := search(heights, low[0], low[1], seen)
		if size > top3[2] {
			top3[2] = size
		}
		if top3[2] > top3[1] {
			top3[1], top3[2] = top3[2], top3[1]
		}
		if top3[1] > top3[0] {
			top3[0], top3[1] = top3[1], top3[0]
		}
	}
	return strconv.Itoa(top3[0] * top3[1] * top3[2]), nil
}

func search(heights [][]int, y, x int, seen [][]bool) int {
	seen[y][x] = true
	size := 1
	for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		yn, xn := y+d[0], x+d[1]
		if yn < 0 || yn >= len(heights) || xn < 0 || xn >= len(heights[yn]) {
			continue
		}
		if seen[yn][xn] {
			continue
		}
		if heights[yn][xn] == 9 {
			continue
		}
		// TODO: Does equal height count to same basin?
		if heights[yn][xn] <= heights[y][x] {
			continue
		}
		size += search(heights, yn, xn, seen)
	}
	return size
}
