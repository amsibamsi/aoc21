package aoc21

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GenFile generates a output file by executing the referenced
// template with the content of the data file passed as a
// string. Data/output filenames are relative to baseDir.
//
// Template inventory:
//   - Package: Base name of baseDir (string)
//   - Data: Contents of the file at baseDir/dataFilename (string)
//
// BUG: This will likely produce incorrect results if any string value
// in the inventory contains backticks and is inserted as-is in the
// template in between backticks.
func GenFile(baseDir, dataFilename, outputFilename, tmplPath string) error {
	var data []byte
	if dataFilename != "" {
		dataFile, err := os.Open(filepath.Join(baseDir, dataFilename))
		if err != nil {
			return err
		}
		defer dataFile.Close()
		data, err = ioutil.ReadAll(dataFile)
		if err != nil {
			return err
		}
	}
	if err := os.MkdirAll(filepath.Dir(filepath.Join(baseDir, outputFilename)), 0777); err != nil {
		return err
	}
	outputFile, err := os.Create(filepath.Join(baseDir, outputFilename))
	if err != nil {
		return err
	}
	defer outputFile.Close()
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	values := struct {
		Package string
		Data    string
	}{
		Package: filepath.Base(baseDir),
		Data:    string(data),
	}
	if err := tmpl.Execute(outputFile, values); err != nil {
		return err
	}
	log.Printf("Created file: %s", outputFile.Name())
	return nil
}
