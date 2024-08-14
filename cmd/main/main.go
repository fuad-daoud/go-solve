package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var s, t string

	fmt.Fscanf(reader, "%s\n", &s)
	fmt.Fscanf(reader, "%s\n", &t)
	index_of_target := 0
	for index := range s {
		if index_of_target < len(t) && (s[index] == t[index_of_target] || s[index] == '?') {
			index_of_target++
		}
	}
	if index_of_target == len(t) {
		fmt.Fprintf(w, "YES\n")
		answer := strings.Builder{}
		index_of_target := 0
		for index := range s {
			if '?' == s[index] && index_of_target < len(t) {
				answer.WriteByte(t[index_of_target])
				index_of_target++
			} else if '?' == s[index] {
				answer.WriteByte('f')
			} else {
				if index_of_target < len(t) && s[index] == t[index_of_target] {
					index_of_target++
				}
				answer.WriteByte(s[index])
			}
		}
		fmt.Fprintf(w, "%s\n", answer.String())

	} else {
		fmt.Fprintf(w, "NO\n")
	}
}
