// Package main runs the puzzle with the sample input.
//
// This file was generated with go:generate.
package main

import (
    "log"

    "github.com/amsibamsi/aoc21"
    "github.com/amsibamsi/aoc21/d03p1"
)

func main() {
    if err := aoc21.RunSolve(d03p1.Solve); err != nil {
        log.Fatal(err)
    }
}