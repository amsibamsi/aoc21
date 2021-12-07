package d07p1

import (
	"math"
	"strconv"
	"strings"
)

func Solve(input string) (string, error) {
	nums := strings.Split(strings.Trim(input, "\n"), ",")
	pos := make([]int, len(nums))
	for i := range nums {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			return "", err
		}
		pos[i] = n
	}
	minPos := math.MaxInt
	maxPos := math.MinInt
	for _, p := range pos {
		if p < minPos {
			minPos = p
		}
		if p > maxPos {
			maxPos = p
		}
	}
	minFuel := math.MaxInt
	for t := minPos; t <= maxPos; t++ {
		fuel := 0
		for _, p := range pos {
			f := p - t
			if f < 0 {
				f = -f
			}
			fuel += f
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return strconv.Itoa(minFuel), nil
}
