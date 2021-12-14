package d13p1

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	dots := [][2]int{}
	for _, line := range aoc21.ToLines(input) {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			fold := strings.Split(strings.Fields(line)[2], "=")
			axis := fold[0]
			pos, err := strconv.Atoi(fold[1])
			if err != nil {
				return "", err
			}
			for i, d := range dots {
				if axis == "x" && d[0] > pos {
					dots[i][0] = 2*pos - d[0]
				}
				if axis == "y" && d[1] > pos {
					dots[i][1] = 2*pos - d[1]
				}
			}
			// Only first fold
			break
		}
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return "", err
		}
		dots = append(dots, [2]int{x, y})
	}
	plot := map[[2]int]bool{}
	for _, d := range dots {
		plot[d] = true
	}
	return strconv.Itoa(len(plot)), nil
}
