package d01p2

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
	for i := 3; i < len(depths); i++ {
		w1 := depths[i-3] + depths[i-2] + depths[i-1]
		w2 := depths[i-2] + depths[i-1] + depths[i]
		if w2 > w1 {
			increases++
		}
	}
	return strconv.Itoa(increases), nil
}
