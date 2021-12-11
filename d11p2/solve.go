package d11p2

import (
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	energy := make([][]int, 0)
	for y, line := range aoc21.ToLines(input) {
		energy = append(energy, make([]int, len(line)))
		for x, char := range line {
			energy[y][x] = int(char) - 48
		}
	}
	height := len(energy)
	width := len(energy[0])
	steps := 0
	for allFlashed := false; !allFlashed; {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				energy[y][x] += 1
			}
		}
		flashed := make([][]bool, height)
		for y := 0; y < height; y++ {
			flashed[y] = make([]bool, width)
		}
		for activity := true; activity; {
			activity = false
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if energy[y][x] > 9 && !flashed[y][x] {
						flashed[y][x] = true
						activity = true
						for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
							yn := y + d[0]
							xn := x + d[1]
							if yn >= 0 && yn < height && xn >= 0 && xn < width {
								energy[yn][xn] += 1
							}
						}
					}
				}
			}
		}
		steps++
		allFlashed = true
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if energy[y][x] > 9 {
					energy[y][x] = 0
				} else {
					allFlashed = false
				}
			}
		}
	}
	return strconv.Itoa(steps), nil
}
