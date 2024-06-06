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

	w := &writer
	var cases int
	fmt.Fscanf(reader, "%d\n", &cases)
	for testCaseNumber := 0; testCaseNumber < cases; testCaseNumber++ {
		solving(reader, w, testCaseNumber)
	}

}

func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var r float64

	fmt.Fscanf(reader, "%v\n", &r)
	var n = r
	answer := 0
	for i := float64(0); i <= n; i++ {
		var x = i * i
		var y = n * n
		var d = math.Sqrt(x + y)
		if r <= d && d < (r+1) {
			answer++
		} else {
			n--
			i--
			x = i * i
			y = n * n
			d = math.Sqrt(x + y)
			if r <= d && d < (r+1) {
				answer++
			}
		}
	}
	for i := n - 1; i > 0; i-- {
		var x = n * n
		var y = i * i
		var d = math.Sqrt(x + y)
		if r <= d && d < (r+1) {
			answer++
		} else {
			n++
			i++
			x = n * n
			y = i * i
			d = math.Sqrt(x + y)
			if r <= d && d < (r+1) {
				answer++
			}
		}
	}
	answer *= 4
	//if n == r && r >= 3 {
	//	answer -= 4
	//}
	fmt.Fprintf(w, "%d\n", answer)
}
