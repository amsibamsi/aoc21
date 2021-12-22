package d21p2

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

const (
	winScore = 21
)

func Solve(input string) (string, error) {
	lines := aoc21.ToLines(input)
	p1, err := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {
		return "", err
	}
	p2, err := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {
		return "", err
	}
	wins := [2]int{0, 0}
	game0 := game{[2]int{p1, p2}, [2]int{0, 0}, 0}
	games := map[game]int{game0: 1}
	for len(games) > 0 {
		ngames := map[game]int{}
		for game, count := range games {
			for _, ngame := range game.play() {
				if w := ngame.winner(); w >= 0 {
					wins[w] += count
					continue
				}
				ngames[*ngame] += count
			}
		}
		games = ngames
	}
	max := wins[0]
	if wins[1] > wins[0] {
		max = wins[1]
	}
	return strconv.Itoa(max), nil
}

type game struct {
	pos   [2]int
	score [2]int
	turn  int
}

func (g *game) play() []*game {
	games := []*game{}
	for i := 0; i < 27; i++ {
		r1 := i/9 + 1
		r2 := (i/3)%3 + 1
		r3 := i%3 + 1
		r := r1 + r2 + r3
		p := (g.pos[g.turn]-1+r)%10 + 1
		s := g.score[g.turn] + p
		npos := g.pos
		npos[g.turn] = p
		nscore := g.score
		nscore[g.turn] = s
		ng := game{npos, nscore, (g.turn + 1) % 2}
		games = append(games, &ng)
	}
	return games
}

func (g *game) winner() int {
	for i, s := range g.score {
		if s >= winScore {
			return i
		}
	}
	return -1
}
