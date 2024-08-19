package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var n string
	fmt.Fscanf(reader, "%s\n", &n)
	if len(n) > 2 && n[0:2] == "10" {
		newN := n[2:len(n)]
		number, _ := strconv.Atoi(newN)
		if number > 1 && n[2] != '0' {
			fmt.Fprintf(w, "YES\n")
		} else {

			fmt.Fprintf(w, "NO\n")
		}
	} else {
		fmt.Fprintf(w, "NO\n")
	}
}
