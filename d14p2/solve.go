package d14p2

import (
	"math"
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	pairs := map[string]int{}
	rules := map[string]byte{}
	counts := map[byte]int{}
	for i, line := range aoc21.ToLines(input) {
		if line == "" {
			continue
		}
		if i == 0 {
			counts[line[0]] += 1
			for j := 1; j < len(line); j++ {
				pairs[line[j-1:j+1]] += 1
				counts[line[j]] += 1
			}
			continue
		}
		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rule[1][0]
	}
	for step := 0; step < 40; step++ {
		diffs := map[string]int{}
		for pair, count := range pairs {
			if insert, ok := rules[pair]; ok {
				pair1 := string(pair[0]) + string(insert)
				pair2 := string(insert) + string(pair[1])
				diffs[pair] -= count
				diffs[pair1] += count
				diffs[pair2] += count
				counts[insert] += count
			}
		}
		for p, c := range diffs {
			pairs[p] += c
		}
	}
	min := math.MaxInt
	max := math.MinInt
	for _, c := range counts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return strconv.Itoa(max - min), nil
}
