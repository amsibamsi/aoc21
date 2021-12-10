package d10p2

import (
	"sort"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	scores := []int{}
	matches := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	for _, line := range aoc21.ToLines(input) {
		stack := []rune{}
		valid := true
		for _, char := range line {
			if match, ok := matches[char]; ok {
				if len(stack) > 0 && match == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					continue
				}
				valid = false
				break
			}
			stack = append(stack, char)

		}
		if valid && len(stack) > 0 {
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = score*5 + points[stack[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return strconv.Itoa(scores[len(scores)/2]), nil
}
