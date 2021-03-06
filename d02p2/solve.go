package d02p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	pos, depth, aim := 0, 0, 0
	for i, line := range aoc21.ToLines(input) {
		cmd := strings.Split(line, " ")
		if len(cmd) != 2 {
			return "", fmt.Errorf("Invalid command on line %d: %s", i, line)
		}
		distance, err := strconv.Atoi(cmd[1])
		if err != nil {
			return "", err
		}
		switch cmd[0] {
		case "forward":
			pos += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			return "", fmt.Errorf("Invalid command on line %d: %s", i, line)
		}
	}
	return strconv.Itoa(pos * depth), nil
}
