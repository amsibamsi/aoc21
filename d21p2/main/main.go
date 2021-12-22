// Package main runs the puzzle with the sample input.
//
// This file was generated with go:generate.
package main

import (
	"log"

	"github.com/amsibamsi/aoc21"
	"github.com/amsibamsi/aoc21/d21p2"
)

func main() {
	if err := aoc21.RunSolve(d21p2.Solve); err != nil {
		log.Fatal(err)
	}
}