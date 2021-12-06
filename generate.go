package aoc21

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
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
func GenFile(baseDir, outputFilename, tmplPath, dataFilename string, overwrite bool) error {
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
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	if !overwrite {
		flags |= os.O_EXCL
	}
	outputFile, err := os.OpenFile(
		filepath.Join(baseDir, outputFilename),
		flags,
		0666,
	)
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
