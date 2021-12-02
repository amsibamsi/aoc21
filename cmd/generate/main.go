// Package main runs all code generation for the project.
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/amsibamsi/aoc21"
)

const (
	baseDir = "../.."
	tmplDir = "../../tmpl"
)

var (
	genFiles = []struct {
		dataFile   string
		outputFile string
		tmplFile   string
	}{
		{"input.txt", "input.go", "input.tmpl"},
		{"output.txt", "output.go", "output.tmpl"},
		{"", "main/main.go", "main.tmpl"},
		{"", "solve_test.go", "test.tmpl"},
		{"", "doc.go", "doc.tmpl"},
	}
)

//go:generate go run main.go
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dirs, err := filepath.Glob(filepath.Join(baseDir, "d[0-9][0-9]p[1-9]"))
	if err != nil {
		return err
	}
	log.Printf("Found directories: %v", dirs)
	for _, d := range dirs {
		for _, g := range genFiles {
			err := aoc21.GenFile(
				d,
				g.dataFile,
				g.outputFile,
				filepath.Join(tmplDir, g.tmplFile),
			)
			if os.IsNotExist(err) {
				log.Print(err)
				continue
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}
