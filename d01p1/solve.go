package d01p1

import (
	"strconv"
)

func Solve(lines []string) (string, error) {
	depths := []int{}
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		depths = append(depths, depth)
	}
	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return strconv.Itoa(increases), nil
}
