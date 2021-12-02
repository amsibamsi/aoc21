package aoc21

import (
	"flag"
	"fmt"
	"io/ioutil"
)

// RunSolve reads the input from a flag argument, calles the specified
// solver and prints the output.
func RunSolve(solver func(input []string) (string, error)) error {
	inputPath := flag.String("input", "input.txt", "File to read input from")
	flag.Parse()
	data, err := ioutil.ReadFile(*inputPath)
	if err != nil {
		return err
	}
	out, err := solver(ToLines(string(data)))
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
