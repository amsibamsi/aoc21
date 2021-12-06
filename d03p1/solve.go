package d03p1

import (
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	lines := aoc21.ToLines(input)
	numLines := len(lines)
	numBits := len(lines[0])
	oneCount := make([]int, numBits)
	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				oneCount[i]++
			}
		}
	}
	gamma := make([]rune, numBits)
	epsilon := make([]rune, numBits)
	for i := range oneCount {
		if oneCount[i] >= numLines/2 {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}
	gammaInt, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		return "", err
	}
	epsilonInt, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(gammaInt * epsilonInt)), nil
}
