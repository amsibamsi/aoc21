package d14p1

import (
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	polymer := ""
	rules := map[string]byte{}
	count := map[byte]int{}
	for i, line := range aoc21.ToLines(input) {
		if line == "" {
			continue
		}
		if i == 0 {
			polymer = line
			continue
		}
		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rule[1][0]
	}
	for i := range polymer {
		count[polymer[i]] += 1
	}
	for step := 0; step < 10; step++ {
		product := bytes.Buffer{}
		product.WriteByte(polymer[0])
		for i := 1; i < len(polymer); i++ {
			if insert, ok := rules[polymer[i-1:i+1]]; ok {
				product.WriteByte(insert)
				count[insert] += 1
			}
			product.WriteByte(polymer[i])
		}
		polymer = product.String()
	}
	min := math.MaxInt
	max := math.MinInt
	for _, c := range count {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return strconv.Itoa(max - min), nil
}
