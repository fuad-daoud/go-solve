package main

import (
	"bufio"
	"os"
	"path/filepath"
)

type TestFile struct {
	input    string
	output   string
	expected string
	name     string
}

func newTestFile(name string) TestFile {
	baseDir := "files"
	return TestFile{
		input:    filepath.Join(baseDir, name, "input.txt"),
		output:   filepath.Join(baseDir, name, "output.txt"),
		expected: filepath.Join(baseDir, name, "expected.txt"),
		name:     name,
	}
}

func inputReader(testFile TestFile) *bufio.Reader {
	input, _ := os.OpenFile(testFile.input, os.O_RDONLY, 0644)
	return bufio.NewReader(input)
}
func outputWriter(testFile TestFile) *bufio.Writer {
	output, _ := os.Create(testFile.output)
	return bufio.NewWriter(output)
}

func ValidateFileEmpty(fileName string) {
	count := 0
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0644)
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		count++
	}
	if count == 0 {
		panic("file is empty " + fileName)
	}
}
