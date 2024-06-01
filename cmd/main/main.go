package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	Solve(reader, *writer)
}

func Solve(reader io.Reader, writer bufio.Writer) {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}(&writer)

	w := &writer
	var cases int
	fmt.Fscanf(reader, "%d\n", &cases)
	for testCaseNumber := 0; testCaseNumber < cases; testCaseNumber++ {
		solving(reader, w, testCaseNumber)
	}

}

func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var s string

	fmt.Fscanf(reader, "%s\n", &s)

	maxStart, maxEnd, answer := getLongestAlreadySortedString(s)
	answer += findRemainingPieces(s, maxStart, maxEnd)

	if len(s) == 1 {
		answer = 1
	}
	fmt.Fprintf(w, "%d\n", max(answer, 1))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findRemainingPieces(s string, maxStart int, maxEnd int) int {
	counter := 0
	for i := 0; i < len(s)-1; i++ {
		if i >= maxStart && i < maxEnd {
			continue
		}
		if s[i] != s[i+1] {
			counter++
		}
	}
	return counter
}

func getLongestAlreadySortedString(s string) (int, int, int) {
	maxStart, maxEnd := -1, -1
	tmpStart, tmpEnd := -1, -1

	startedSeeingZeros, startedSeeingOnes := false, false
	for index, value := range s {
		if value == '0' && !startedSeeingZeros {
			startedSeeingZeros = true
			tmpStart = index
		}
		if value == '1' && startedSeeingZeros {
			startedSeeingOnes = true
		}
		if value == '0' && startedSeeingOnes {
			tmpEnd = index - 1
			startedSeeingOnes = false

			if tmpEnd-tmpStart > maxEnd-maxStart {
				maxEnd = tmpEnd
				maxStart = tmpStart
			}
			tmpEnd = -1
			tmpStart = index
		}
	}
	if maxEnd == -1 && tmpStart != -1 {
		maxStart = tmpStart
		maxEnd = len(s)
	}
	if maxStart != -1 && maxEnd != -1 {
		return maxStart, maxEnd, 1
	}
	return maxStart, maxEnd, 0
}
