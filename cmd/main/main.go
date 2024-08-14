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

// took a look at kilani's solution
func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var a1, a2, b1, b2 int

	fmt.Fscanf(reader, "%v %v %v %v\n", &a1, &a2, &b1, &b2)
	answer := 0
	if get(a1, b1)+get(a2, b2) > 0 {
		answer += 2
	}
	if get(a2, b1)+get(a1, b2) > 0 {
		answer += 2
	}
	fmt.Fprintf(w, "%v\n", answer)

}
func get(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
