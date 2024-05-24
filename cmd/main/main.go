package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	r := &writer
	var cases int
	fmt.Fscanf(reader, "%d\n", &cases)
	for testCaseNumber := 0; testCaseNumber < cases; testCaseNumber++ {
		solving(reader, r, testCaseNumber)
	}

}

func solving(reader io.Reader, r *bufio.Writer, testCaseNumber int) {
	var a, b float64

	fmt.Fscanf(reader, "%f %f\n", &a, &b)
	fmt.Fprintln(r, math.Min(a, b), math.Max(a, b))
}
