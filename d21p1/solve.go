package d21p1

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	lines := aoc21.ToLines(input)
	pos := [2]int{}
	var err error
	pos[0], err = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {
		return "", err
	}
	pos[1], err = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {
		return "", err
	}
	dice := 1
	rolls := 0
	score := [2]int{}
	for score[1] < 1000 {
		for i := 0; i < 3; i++ {
			pos[0] = (pos[0]+dice-1)%10 + 1
			dice = dice%100 + 1
			rolls++
		}
		score[0] += pos[0]
		pos[0], pos[1] = pos[1], pos[0]
		score[0], score[1] = score[1], score[0]
	}
	return strconv.Itoa(score[0] * rolls), nil
}
