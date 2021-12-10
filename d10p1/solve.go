package d10p1

import (
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	score := 0
	matches := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, line := range aoc21.ToLines(input) {
		stack := []rune{}
		for _, char := range line {
			if match, ok := matches[char]; ok {
				if len(stack) > 0 && match == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					continue
				}
				score += scores[char]
				break
			}
			stack = append(stack, char)

		}
	}
	return strconv.Itoa(score), nil
}
