package d06p1

import (
	"strconv"
	"strings"
)

func Solve(lines []string) (string, error) {
	fish := make([]int, 0)
	for _, num := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return "", err
		}
		fish = append(fish, n)
	}
	for day := 0; day < 80; day++ {
		for i := range fish {
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
				continue
			}
			fish[i] -= 1
		}
	}
	return strconv.Itoa(len(fish)), nil
}
