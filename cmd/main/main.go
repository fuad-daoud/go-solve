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

const MAX = 200_000_7

var a, sum [MAX]int

func steps(x int) int {
	counter := 0
	for {
		if x < 1 {
			break
		}
		x /= 3
		counter++
	}
	return counter
}

func pre_solve() {
	a = [MAX]int{}
	sum = [MAX]int{}
	for index := 1; index < MAX; index++ {
		a[index] = steps(index)
		sum[index] = sum[index-1] + a[index]
	}
}

// lookuped the editorial
func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var left, right int
	fmt.Fscanf(reader, "%v %v\n", &left, &right)
	answer := sum[right] - sum[left-1] + a[left]
	fmt.Fprintf(w, "%v\n", answer)
}
