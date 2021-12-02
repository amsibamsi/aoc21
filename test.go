package aoc21

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T, solver func(lines []string) (string, error), input, output string) {
	want := strings.Trim(output, "\n")
	got, err := solver(ToLines(input))
	if err == nil {
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	} else {
		t.Errorf("Unexpected error: %v", err)
	}
}
