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
	var s string

	fmt.Fscanf(reader, "%s\n", &s)
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch] = freq[ch] + 1
	}

	keys := make([]rune, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	if len(keys) == 1 {
		fmt.Fprintln(w, "NO")
		return
	}
	fmt.Fprintln(w, "YES")
	var strBuilder strings.Builder

	for key := range freq {
		for index := 0; index < freq[key]; index++ {
			strBuilder.WriteString(string(key))
		}
	}
	answer := strBuilder.String()
	if strings.Compare(answer, s) == 0 {
		answer = reverse(answer)
	}
	fmt.Fprintf(w, "%s\n", answer)

}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
