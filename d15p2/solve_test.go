// This file was generated with go:generate.
package d15p2

import (
	"testing"

	"github.com/amsibamsi/aoc21"
)

func TestSolve(t *testing.T) {
	if testing.Short() {
		t.Skip("Very slow, ~400 secs")
	}
	aoc21.TestSolve(t, Solve, Input, Output)
}