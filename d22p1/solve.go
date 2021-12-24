package d22p1

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	r := newReactor()
	for _, line := range aoc21.ToLines(input) {
		parts := strings.Split(line, " ")
		power := true
		if parts[0] == "off" {
			power = false
		}
		cuboid := map[string][2]int{}
		for _, coord := range strings.Split(parts[1], ",") {
			spec := strings.Split(coord, "=")
			axis := spec[0]
			minmax := strings.Split(spec[1], "..")
			min, err := strconv.Atoi(minmax[0])
			if err != nil {
				return "", err
			}
			max, err := strconv.Atoi(minmax[1])
			if err != nil {
				return "", err
			}
			if min < -50 {
				min = -50
			}
			if max > 50 {
				max = 50
			}
			cuboid[axis] = [2]int{min, max}
		}
		for x := cuboid["x"][0]; x <= cuboid["x"][1]; x++ {
			for y := cuboid["y"][0]; y <= cuboid["y"][1]; y++ {
				for z := cuboid["z"][0]; z <= cuboid["z"][1]; z++ {
					r.set(x, y, z, power)
				}
			}
		}
	}
	return strconv.Itoa(r.countOn()), nil
}

type reactor struct {
	cubes map[[3]int]bool
}

func newReactor() *reactor {
	return &reactor{map[[3]int]bool{}}
}

func (r *reactor) set(x, y, z int, power bool) {
	if power {
		r.cubes[[3]int{x, y, z}] = true
		return
	}
	delete(r.cubes, [3]int{x, y, z})
}

func (r *reactor) countOn() int {
	return len(r.cubes)
}
