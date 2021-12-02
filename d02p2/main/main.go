// Package main runs the puzzle with the sample input.
package main

import (
    "log"

    "github.com/amsibamsi/aoc21"
    "github.com/amsibamsi/aoc21/d02p2"
)

func main() {
    if err := aoc21.RunSolve(d02p2.Solve); err != nil {
        log.Fatal(err)
    }
}