// Package main solves AoC 2021 Day 1.
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
			log.Printf("Line %d is not a number: %v", i, err)
			continue
		}
		depths = append(depths, depth)
	}
	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	fmt.Println(increases)
	return nil
}
