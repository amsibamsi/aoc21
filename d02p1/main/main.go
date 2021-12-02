// Package main runs the puzzle with the sample input.
package main

import (
    "log"

    "github.com/amsibamsi/aoc21"
    "github.com/amsibamsi/aoc21/d02p1"
)

func main() {
    if err := aoc21.RunSolve(d02p1.Solve); err != nil {
        log.Fatal(err)
    }
}