package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"testing"
)

func Test(t *testing.T) {
	testFiles := []TestFile{
		newTestFile("basic"),
	}
	fmt.Printf("%+v\n", testFiles)

	for _, testFile := range testFiles {
		t.Run(testFile.name, func(t *testing.T) {
			fmt.Println("Starting test", testFile.name)
			ValidateFileEmpty(testFile.input)
			ValidateFileEmpty(testFile.expected)

			solver := Solver{reader: inputReader(testFile), writer: outputWriter(testFile)}
			solver.Solve()
			validateOutput(testFile.output, testFile.expected, t)
		})
	}
}

func validateOutput(output, expected string, t *testing.T) {
	outputFileLines := readLines(output)
	expectedFileLines := readLines(expected)

	n := getMin(outputFileLines, expectedFileLines)
	for i := 0; i < n; i++ {
		if outputFileLines[i] != expectedFileLines[i] {
			_, _ = fmt.Printf("failed at line %d: expected %q, got %q\n", i+1, expectedFileLines[i], outputFileLines[i])
			t.Fail()
		}
	}
	if len(outputFileLines) != len(expectedFileLines) {
		_, _ = fmt.Printf("file line count mismatch, outputFileLines: %d, expectedFileLines: %d\n", len(outputFileLines), len(expectedFileLines))
		t.Fail()
	}
}

func readLines(fileName string) (lines []string) {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0644)
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	fmt.Printf("lines for file:\n%v %q\n", fileName, lines)
	return lines
}

func getMin(first, second []string) int {
	return int(math.Min(float64(len(first)), float64(len(second))))
}
