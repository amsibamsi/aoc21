package d06p2

import (
	"strconv"
	"strings"
)

func Solve(input string) (string, error) {
	pop := [9]int{}
	for _, num := range strings.Split(strings.Trim(input, "\n"), ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return "", err
		}
		pop[n] += 1
	}
	for day := 0; day < 256; day++ {
		popZero := pop[0]
		for i := 0; i < 6; i++ {
			pop[i] = pop[i+1]
		}
		pop[6] = pop[7] + popZero
		pop[7] = pop[8]
		pop[8] = popZero

	}
	total := 0
	for _, n := range pop {
		total += n
	}
	return strconv.Itoa(total), nil
}
