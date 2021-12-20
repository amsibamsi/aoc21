package d18p2

import (
	"fmt"
	"math"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	max := math.MinInt
	numbers := []*number{}
	for _, line := range aoc21.ToLines(input) {
		n, _ := newNumber(line, 0)
		numbers = append(numbers, n)
	}
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 == n2 {
				continue
			}
			sum := n1.copy()
			sum.add(n2.copy())
			for changed := true; changed; {
				es := explodeStatus{}
				sum.explode(0, &es)
				changed = es.exploded
				if !changed {
					changed = sum.split(0)
				}
			}
			m := sum.magnitude()
			if m > max {
				max = m
			}
		}
	}
	return strconv.Itoa(max), nil
}

type number struct {
	value       int
	left, right *number
}

func newNumber(s string, p int) (*number, int) {
	if s[p] == '[' {
		l, end := newNumber(s, p+1)
		r, end := newNumber(s, end+2)
		return &number{-1, l, r}, end + 1
	}
	num := int(s[p]) - 48
	return &number{num, nil, nil}, p
}

func (n *number) add(m *number) {
	left := &number{n.value, n.left, n.right}
	n.left = left
	n.right = m
	n.value = -1
}

func (n *number) magnitude() int {
	if n.value >= 0 {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *number) copy() *number {
	if n.value >= 0 {
		return &number{n.value, nil, nil}
	}
	return &number{-1, n.left.copy(), n.right.copy()}
}

func (n number) String() string {
	if n.value >= 0 {
		return strconv.Itoa(n.value)
	}
	return fmt.Sprintf("[%v,%v]", n.left.String(), n.right.String())
}

type explodeStatus struct {
	leftVal  *number
	exploded bool
	addRight int
}

func (n *number) explode(level int, es *explodeStatus) {
	if n == nil {
		return
	}
	if es.exploded && es.addRight < 0 {
		return
	}
	if n.value >= 0 {
		if es.exploded {
			n.value += es.addRight
			es.addRight = -1
			return
		}
		es.leftVal = n
		return
	}
	if !es.exploded && level >= 4 && n.left.value >= 0 && n.right.value >= 0 {
		if es.leftVal != nil {
			es.leftVal.value += n.left.value
		}
		es.exploded = true
		es.addRight = n.right.value
		n.left = nil
		n.right = nil
		n.value = 0
		return
	}
	n.left.explode(level+1, es)
	n.right.explode(level+1, es)
}

func (n *number) split(level int) bool {
	if n == nil {
		return false
	}
	if n.value >= 10 {
		n.left = &number{n.value / 2, nil, nil}
		n.right = &number{(n.value + 1) / 2, nil, nil}
		n.value = -1
		return true
	}
	if n.value >= 0 {
		return false
	}
	if n.left.split(level + 1) {
		return true
	}
	return n.right.split(level + 1)
}
