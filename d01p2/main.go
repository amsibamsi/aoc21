// Package main solves day 1, part 2.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	depths := []int{}
	for i, line := range lines {
		if line == "" {
			log.Printf("Line %d is empty", i)
			continue
		}
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Line %d is not an integer: %v", i, err)
			continue
		}
		depths = append(depths, depth)
	}
	increases := 0
	for i := 3; i < len(depths); i++ {
		w1 := depths[i-3] + depths[i-2] + depths[i-1]
		w2 := depths[i-2] + depths[i-1] + depths[i]
		if w2 > w1 {
			increases++
		}
	}
	fmt.Println(increases)
	return nil
}
