package d22p2

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	var cuboids []*cuboid
	for _, line := range aoc21.ToLines(input) {
		parts := strings.Split(line, " ")
		power := 1
		if parts[0] == "off" {
			power = -1
		}
		coords := []int{}
		for _, coord := range strings.Split(parts[1], ",") {
			minmax := strings.Split(strings.Split(coord, "=")[1], "..")
			min, err := strconv.Atoi(minmax[0])
			if err != nil {
				return "", err
			}
			max, err := strconv.Atoi(minmax[1])
			if err != nil {
				return "", err
			}
			coords = append(coords, min, max)
		}
		ncuboid := newCuboid(
			coords[0],
			coords[1],
			coords[2],
			coords[3],
			coords[4],
			coords[5],
			power,
		)
		if len(cuboids) == 0 && power == 1 {
			cuboids = []*cuboid{ncuboid}
			continue
		}
		ncuboids := make([]*cuboid, len(cuboids))
		copy(ncuboids, cuboids)
		if power == 1 {
			ncuboids = append(ncuboids, ncuboid)
		}
		for _, c := range cuboids {
			is := intersectSub(c, ncuboid)
			if is != nil {
				ncuboids = append(ncuboids, is)
			}
		}
		cuboids = ncuboids
	}
	count := 0
	for _, c := range cuboids {
		count += c.countCubes()
	}
	return strconv.Itoa(count), nil
}

type cuboid struct {
	x1, x2, y1, y2, z1, z2, addSub int
}

func newCuboid(x1, x2, y1, y2, z1, z2, addSub int) *cuboid {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if z2 < z1 {
		z1, z2 = z2, z1
	}
	return &cuboid{x1, x2, y1, y2, z1, z2, addSub}
}

func (c *cuboid) countCubes() int {
	return c.addSub *
		((c.x2 - c.x1) + 1) *
		((c.y2 - c.y1) + 1) *
		((c.z2 - c.z1) + 1)
}

func intersectSub(a, b *cuboid) *cuboid {
	if a.x1 > b.x2 || a.x2 < b.x1 ||
		a.y1 > b.y2 || a.y2 < b.y1 ||
		a.z1 > b.z2 || a.z2 < b.z1 {
		return nil
	}
	c := cuboid{
		max(a.x1, b.x1), min(a.x2, b.x2),
		max(a.y1, b.y1), min(a.y2, b.y2),
		max(a.z1, b.z1), min(a.z2, b.z2),
		b.addSub,
	}
	if a.addSub == b.addSub {
		c.addSub = -c.addSub
	}
	return &c
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// func unionDistinct(a, b *cuboid) []*cuboid {
// 	if a.x1 > b.x2 || a.x2 < b.x1 ||
// 		a.y1 > b.y2 || a.y2 < b.y1 ||
// 		a.z1 > b.z2 || a.z2 < b.z1 {
// 		return []*cuboid{a, b}
// 	}
// 	is := cuboid{power: b.power}
// 	set := []*cuboid{&is}
// 	nCuboid := func(x1, x2, y1, y2, z1, z2 int, power bool) {
// 		set = append(set, newCuboid(x1, x2, y1, y2, z1, z2, power))
// 	}
// 	firstLast := func(a, b *cuboid, n, m int) (*cuboid, *cuboid, bool) {
// 		if n < m {
// 			return a, b, false
// 		}
// 		if m < n {
// 			return b, a, false
// 		}
// 		return a, b, true
// 	}
// 	var f, l *cuboid
// 	var e bool
// 	f, l, e = firstLast(a, b, a.x1, b.x1)
// 	is.x1 = l.x1
// 	if !e {
// 		nCuboid(f.x1, l.x1-1, f.y1, f.y2, f.z1, f.z2, f.power)
// 	}
// 	f, l, e = firstLast(a, b, a.x2, b.x2)
// 	is.x2 = f.x2
// 	if !e {
// 		nCuboid(f.x2+1, l.x2, l.y1, l.y2, l.z1, l.z2, l.power)
// 	}
// 	f, l, e = firstLast(a, b, a.y1, b.y1)
// 	is.y1 = l.y1
// 	if !e {
// 		nCuboid(is.x1, is.x2, f.y1, l.y1-1, f.z1, f.z2, f.power)
// 	}
// 	f, l, e = firstLast(a, b, a.y2, b.y2)
// 	is.y2 = f.y2
// 	if !e {
// 		nCuboid(is.x1, is.x2, f.y2+1, l.y2, l.z1, l.z2, l.power)
// 	}
// 	f, l, e = firstLast(a, b, a.z1, b.z1)
// 	is.z1 = l.z1
// 	if !e {
// 		nCuboid(is.x1, is.x2, is.y1, is.y2, f.z1, l.z1-1, f.power)
// 	}
// 	f, l, e = firstLast(a, b, a.z2, b.z2)
// 	is.z2 = f.z2
// 	if !e {
// 		nCuboid(is.x1, is.x2, is.y1, is.y2, f.z2+1, l.z2, l.power)
// 	}
// 	return set
// }
