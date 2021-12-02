// Package main solves day 2, part 2.
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
	pos, depth, aim := 0, 0, 0
	for i, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		if line == "" {
			log.Printf("Line %d is empty", i)
			continue
		}
		cmd := strings.Split(line, " ")
		if len(cmd) != 2 {
			log.Printf("Invalid command on line %d: %s", i, line)
			continue
		}
		distance, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Printf("Invalid distance on line %d: %v", i, err)
		}
		switch cmd[0] {
		case "forward":
			pos += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			log.Printf("Invalid command on line %d: %s", i, line)
		}
	}
	fmt.Println(pos * depth)
	return nil
}
