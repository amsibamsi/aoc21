// Package main runs the puzzle with the sample input.
package main

import (
    "log"

    "github.com/amsibamsi/aoc21"
    "github.com/amsibamsi/aoc21/d01p1"
)

func main() {
    if err := aoc21.RunSolve(d01p1.Solve); err != nil {
        log.Fatal(err)
    }
}