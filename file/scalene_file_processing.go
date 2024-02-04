package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

var fileMap = map[string]bool{
	"yaml": true,
	"json": true,
	"yml":  true,
}

func recurFile(f fs.DirEntry) {
	if f.IsDir() {
		files, err := os.ReadDir("/go-learning")

		if err != nil {
			return
		}

		for _, file := range files {
			recurFile(file)
		}
	}

	name := f.Name()
	values := strings.Split(name, ".")
	if len(values) > 1 {
		if fileMap[values[1]] {
			fmt.Println(name)
		}
	}
}

// WAP to print all file name recursively from directory whose extension is ending with yaml, json, yml
func scaleneInterviewFileProcessing() {
	files, err := os.ReadDir("/go-learning")

	if err != nil {
		return
	}

	for _, file := range files {
		recurFile(file)
	}

	// from here below code is printing words whose starting with A to 9 from file
	file, _ := os.Open("fileName")

	scanner := bufio.NewScanner(file)

	for i, val := range scanner.Text() {
		match, _ := regexp.MatchString("[A-9]", string(val))
		if !match {
			fmt.Println(i + 1)
		}
	}
}
