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
	var a, b, c, d int

	fmt.Fscanf(reader, "%d %d %d %d\n", &a, &b, &c, &d)
	counter := 0
	for i := min(a, b); i <= max(a, b); i++ {
		if i == c || i == d {
			counter++
		}
	}
	if counter == 1 {
		fmt.Fprintln(w, "YES")
		return
	}
	counter = 0
	for i := min(c, d); i <= max(c, d); i++ {
		if i == a || i == b {
			counter++
		}
	}
	if counter == 1 {
		fmt.Fprintln(w, "YES")
		return
	}
	fmt.Fprintln(w, "NO")

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
