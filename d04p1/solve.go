package d04p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	lines := aoc21.ToLines(input)
	draws := strings.Split(lines[0], ",")
	boardSrc := strings.Join(lines[2:], "\n")
	var boards []*board
	for _, b := range strings.Split(boardSrc, "\n\n") {
		boards = append(boards, newBoard(b))
	}
	var lastDraw string
	var winner *board
	for _, d := range draws {
		for _, b := range boards {
			if b.draw(d) {
				lastDraw = d
				winner = b
				goto winner
			}
		}
	}
	return "", fmt.Errorf("No winner")
winner:
	drawScore, err := strconv.Atoi(lastDraw)
	if err != nil {
		return "", err
	}
	score, err := winner.score()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(score * drawScore), nil
}

type board struct {
	numbers        [][]string
	crossed        [][]bool
	rowsCrossed    []int
	columnsCrossed []int
	bingo          bool
}

func newBoard(s string) *board {
	b := board{}
	b.numbers = make([][]string, 0)
	b.crossed = make([][]bool, 0)
	for _, line := range strings.Split(s, "\n") {
		nums := strings.Fields(line)
		b.numbers = append(b.numbers, nums)
		b.crossed = append(b.crossed, make([]bool, len(nums)))
	}
	b.rowsCrossed = make([]int, len(b.numbers))
	b.columnsCrossed = make([]int, len(b.numbers[0]))
	return &b
}

func (b *board) draw(s string) bool {
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if b.numbers[i][j] == s && !b.crossed[i][j] {
				b.crossed[i][j] = true
				b.rowsCrossed[i]++
				b.columnsCrossed[j]++
				if b.rowsCrossed[i] == len(b.numbers) || b.columnsCrossed[j] == len(b.numbers[0]) {
					b.bingo = true
				}
			}
		}
	}
	return b.bingo
}

func (b *board) score() (int, error) {
	score := 0
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if !b.crossed[i][j] {
				n, err := strconv.Atoi(b.numbers[i][j])
				if err != nil {
					return -1, err
				}
				score += n
			}
		}
	}
	return score, nil
}
