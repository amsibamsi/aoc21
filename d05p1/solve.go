package d05p1

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	d := diagram{}
	for _, line := range aoc21.ToLines(input) {
		points := strings.Split(line, " -> ")
		p0, err := parseInts(strings.Split(points[0], ","))
		if err != nil {
			return "", err
		}
		p1, err := parseInts(strings.Split(points[1], ","))
		if err != nil {
			return "", err
		}
		if p0[0] != p1[0] && p0[1] != p1[1] {
			continue
		}
		dx, dy := 1, 1
		if p1[0] < p0[0] {
			dx = -1
		}
		if p1[1] < p0[1] {
			dy = -1
		}
		for x := p0[0]; x != p1[0]+dx; x += dx {
			for y := p0[1]; y != p1[1]+dy; y += dy {
				d.cover(x, y)
			}
		}
	}
	return strconv.Itoa(d.count(2)), nil
}

func parseInts(s []string) ([]int, error) {
	ints := make([]int, len(s))
	for i, text := range s {
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

type diagram struct {
	vents [][]int
}

func (d *diagram) cover(x, y int) {
	if x >= len(d.vents) {
		for i := len(d.vents); i <= x; i++ {
			d.vents = append(d.vents, nil)
		}
	}
	if d.vents[x] == nil {
		d.vents[x] = make([]int, y+1)
	}
	if y >= len(d.vents[x]) {
		for i := len(d.vents[x]); i <= y; i++ {
			d.vents[x] = append(d.vents[x], 0)
		}
	}
	d.vents[x][y] += 1
}

func (d *diagram) count(min int) int {
	sum := 0
	for _, row := range d.vents {
		for _, count := range row {
			if count >= min {
				sum++
			}
		}
	}
	return sum
}
