package d18p1

import (
	"fmt"
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	var total *number
	for i, line := range aoc21.ToLines(input) {
		n, _ := newNumber(line, 0)
		if i == 0 {
			total = n
			continue
		}
		total.add(n)
		for changed := true; changed; {
			es := explodeStatus{}
			total.explode(0, &es)
			changed = es.exploded
			if !changed {
				changed = total.split(0)
			}
		}
	}
	return strconv.Itoa(total.magnitude()), nil
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
