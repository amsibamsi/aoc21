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
		outputFile string
		tmplFile   string
		dataFile   string
		overwrite  bool
	}{
		{"input.txt", "input.txt.tmpl", "", false},
		{"output.txt", "output.txt.tmpl", "", false},
		{"solve.go", "solve.go.tmpl", "", false},
		{"main/main.go", "main.go.tmpl", "", true},
		{"solve_test.go", "solve_test.go.tmpl", "", true},
		{"doc.go", "doc.go.tmpl", "", true},
		{"input.go", "input.go.tmpl", "input.txt", true},
		{"output.go", "output.go.tmpl", "output.txt", true},
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
				g.outputFile,
				filepath.Join(tmplDir, g.tmplFile),
				g.dataFile,
				g.overwrite,
			)
			if os.IsNotExist(err) {
				log.Print(err)
				continue
			}
			if os.IsExist(err) {
				continue
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}
