package d03p2

import (
	"strconv"
)

func Solve(lines []string) (string, error) {
	numLines := len(lines)
	numBits := len(lines[0])
	candidates := make([]string, numLines)
	copy(candidates, lines)
	for pos := 0; pos < numBits && len(candidates) > 1; pos++ {
		oneCount := 0
		for _, line := range candidates {
			if line[pos] == '1' {
				oneCount++
			}
		}
		keep := byte('1')
		if oneCount < len(candidates)-oneCount {
			keep = '0'
		}
		n := 0
		for _, line := range candidates {
			if line[pos] == keep {
				candidates[n] = line
				n++
			}
		}
		candidates = candidates[:n]
	}
	oxygen, err := strconv.ParseInt(candidates[0], 2, 64)
	if err != nil {
		return "", err
	}
	candidates = make([]string, numLines)
	copy(candidates, lines)
	for pos := 0; pos < numBits && len(candidates) > 1; pos++ {
		oneCount := 0
		for _, line := range candidates {
			if line[pos] == '1' {
				oneCount++
			}
		}
		keep := byte('0')
		if oneCount < len(candidates)-oneCount {
			keep = '1'
		}
		n := 0
		for _, line := range candidates {
			if line[pos] == keep {
				candidates[n] = line
				n++
			}
		}
		candidates = candidates[:n]
	}
	co2, err := strconv.ParseInt(candidates[0], 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(oxygen) * int(co2)), err
}
