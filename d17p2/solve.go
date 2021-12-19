package d17p2

import (
	"strconv"
	"strings"
)

func Solve(input string) (string, error) {
	coords := strings.Split(strings.Split(strings.Trim(input, "\n"), ": ")[1], ", ")
	xcoords := strings.Split(strings.Split(coords[0], "=")[1], "..")
	ycoords := strings.Split(strings.Split(coords[1], "=")[1], "..")
	xmin, err := strconv.Atoi(xcoords[0])
	if err != nil {
		return "", err
	}
	xmax, err := strconv.Atoi(xcoords[1])
	if err != nil {
		return "", err
	}
	if xmin > xmax {
		xmin, xmax = xmax, xmin
	}
	ymin, err := strconv.Atoi(ycoords[0])
	if err != nil {
		return "", err
	}
	ymax, err := strconv.Atoi(ycoords[1])
	if err != nil {
		return "", err
	}
	if ymin > ymax {
		ymin, ymax = ymax, ymin
	}
	vxmin := 0
	vxmax := xmax
	vymin := 1
	if ymin < 0 {
		vymin = ymin
	}
	vymax := ymax
	if ymax < 0 {
		vymax = -ymin - 1
	}
	count := 0
	for vx := vxmin; vx <= vxmax; vx++ {
		for vy := vymin; vy <= vymax; vy++ {
			if launch(vx, vy, xmin, xmax, ymin, ymax) {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}

func launch(vx, vy, xmin, xmax, ymin, ymax int) bool {
	x, y := 0, 0
	for {
		x += vx
		y += vy
		if vx > 0 {
			vx--
		}
		vy--
		if x > xmax {
			return false
		}
		if x < xmin && vx == 0 {
			return false
		}
		if y < ymin && vy <= 0 {
			return false
		}
		if x >= xmin && x <= xmax && y >= ymin && y <= ymax {
			return true
		}
	}
}
