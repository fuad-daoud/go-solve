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

func solving(r io.Reader, w *bufio.Writer, testCaseNumber int) {
	var n int
	fmt.Fscanf(r, "%d\n", &n)

	nums := make([]int, n)
	for index := 0; index < n; index++ {
		fmt.Fscanf(r, "%d", &nums[index])
		if index != n-1 {
			fmt.Fscanf(r, " ")
		}
	}
	fmt.Fscanf(r, "\n")
	visited := make(map[int]bool)

	counter := 0

	for index := 0; index < n; index++ {
		visited[nums[index]] = true
		if !(visited[nums[index]-1] || visited[nums[index]+1]) {
			counter++
		}
	}

	if counter > 1 {
		fmt.Fprint(w, "NO\n")
	} else {
		fmt.Fprint(w, "YES\n")
	}
}
