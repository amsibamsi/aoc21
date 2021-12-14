package d13p2

import (
	"bytes"
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
			continue
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
	merged := map[[2]int]bool{}
	for _, d := range dots {
		merged[d] = true
	}
	max := [2]int{0, 0}
	for _, d := range dots {
		if d[0] > max[0] {
			max[0] = d[0]
		}
		if d[1] > max[1] {
			max[1] = d[1]
		}
	}
	plot := bytes.Buffer{}
	for y := 0; y <= max[1]; y++ {
		for x := 0; x <= max[0]; x++ {
			char := " "
			if _, ok := merged[[2]int{x, y}]; ok {
				char = "#"
			}
			plot.WriteString(char)
		}
		if y != max[1] {
			plot.WriteString("\n")
		}
	}
	return plot.String(), nil
}
