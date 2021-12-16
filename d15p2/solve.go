package d15p2

import (
	"container/heap"
	"math"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	levels := map[[2]int]int{}
	lines := aoc21.ToLines(input)
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, level := range line {
			num, err := strconv.Atoi(string(level))
			if err != nil {
				return "", err
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					l := (num+i+j-1)%9 + 1
					levels[[2]int{x + i*width, y + j*height}] = l
				}
			}
		}
	}
	start := [2]int{0, 0}
	end := [2]int{5*width - 1, 5*height - 1}
	risks := map[[2]int]int{}
	max := math.MaxInt
	nodes := newGraph()
	heap.Push(nodes, &node{start, -levels[start]})
	for n := heap.Pop(nodes).(*node); n != nil; n = heap.Pop(nodes).(*node) {
		if _, ok := levels[n.pos]; !ok {
			continue
		}
		risk := n.risk + levels[n.pos]
		if risk >= max {
			continue
		}
		if r, ok := risks[n.pos]; ok && risk >= r {
			continue
		}
		risks[n.pos] = risk
		if n.pos == end {
			if risk < max {
				max = risk
			}
			continue
		}
		for _, d := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			next := [2]int{n.pos[0] + d[0], n.pos[1] + d[1]}
			if next == n.pos {
				continue
			}
			heap.Push(nodes, &node{next, risk})
		}
	}
	return strconv.Itoa(risks[end]), nil
}

type node struct {
	pos  [2]int
	risk int
}

type graph struct {
	data []*node
}

func newGraph() *graph {
	return &graph{[]*node{}}
}

func (g *graph) Len() int {
	return len(g.data)
}

func (g *graph) Less(i, j int) bool {
	return g.data[i].risk < g.data[j].risk
}

func (g *graph) Swap(i, j int) {
	if len(g.data) == 0 {
		return
	}
	g.data[i], g.data[j] = g.data[j], g.data[i]
}

func (g *graph) Push(n interface{}) {
	g.data = append(g.data, n.(*node))
}

func (g *graph) Pop() interface{} {
	if len(g.data) == 0 {
		return (*node)(nil)
	}
	p := g.data[len(g.data)-1]
	g.data = g.data[:len(g.data)-1]
	return p
}
