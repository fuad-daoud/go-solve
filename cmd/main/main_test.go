package main

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	baseDir := "files"
	tests := []test{
		newTest("basic"),
	}
	for _, test := range tests {
		test.input = filepath.Join(baseDir, test.name+"-basic-input.txt")
		test.output = filepath.Join(baseDir, test.name+"-output.txt")
		test.expected = filepath.Join(baseDir, test.name+"-basic-expected.txt")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, _ := os.OpenFile(tt.input, os.O_RDONLY, 0644)
			reader := bufio.NewReader(input)
			output, _ := os.Create(tt.output)
			writer := bufio.NewWriter(output)
			Solve(reader, *writer)
			equals, _ := FileCmp(tt.output, tt.output)
			if !equals {
				t.FailNow()
			}
		})
	}
}

type test struct {
	input    string
	output   string
	expected string
	name     string
}

func newTest(name string) test {
	return test{
		input:    name,
		output:   name,
		expected: name,
		name:     name,
	}
}

func FileCmp(file1, file2 string) (same bool, err error) {

	stat1, err := os.Stat(file1)
	if err != nil {
		return false, err
	}

	stat2, err := os.Stat(file2)
	if err != nil {
		return false, err
	}

	if os.SameFile(stat1, stat2) {
		return true, nil
	}

	if stat1.Size() != stat2.Size() {
		return false, nil
	}
	return true, nil
}
