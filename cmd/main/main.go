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
	var n, s, m int

	fmt.Fscanf(reader, "%v %v %v\n", &n, &s, &m)
	yes := false
	pre_left := 0
	for i := 0; i < n; i++ {
		var left, right int
		fmt.Fscanf(reader, "%v %v\n", &left, &right)
		if left-pre_left >= s {
			yes = true
		}
		pre_left = right
	}
	if m-pre_left >= s {
		yes = true
	}
	if yes {
		fmt.Fprintf(w, "YES\n")
	} else {
		fmt.Fprintf(w, "NO\n")
	}
}
