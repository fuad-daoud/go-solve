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
	pre_solve()
	for testCaseNumber := 0; testCaseNumber < cases; testCaseNumber++ {
		solving(reader, w, testCaseNumber)
	}

}

func pre_solve() {
}

// lookuped the editorial
func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var n int
	fmt.Fscanf(reader, "%v\n", &n)

	answer := (n / 4) + ((n / 2) % 2)

	fmt.Fprintf(w, "%v\n", answer)
}
