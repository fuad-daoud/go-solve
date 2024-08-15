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

func solving(reader io.Reader, w *bufio.Writer, testCaseNumber int) {
	var n, q int
	fmt.Fscanf(reader, "%v %v\n", &n, &q)

	var a, b string
	fmt.Fscanf(reader, "%v\n%v\n", &a, &b)

	a_freqs := make([][]int, n+1)
	b_freqs := make([][]int, n+1)

	a_freqs[0] = make([]int, 26)
	b_freqs[0] = make([]int, 26)
	for index := 1; index <= n; index++ {
		a_freqs[index] = make([]int, 26)
		b_freqs[index] = make([]int, 26)

		for c_index := 0; c_index < 26; c_index++ {
			a_freqs[index][c_index] = a_freqs[index-1][c_index]
			b_freqs[index][c_index] = b_freqs[index-1][c_index]
		}
		a_freqs[index][a[index-1]-'a']++
		b_freqs[index][b[index-1]-'a']++
	}
	for query := 0; query < q; query++ {
		var left, right int
		fmt.Fscanf(reader, "%v %v\n", &left, &right)

		answer := 0

		for index := 0; index < 26; index++ {
			a_char_freqs := a_freqs[right][index] - a_freqs[left-1][index]

			b_char_freqs := b_freqs[right][index] - b_freqs[left-1][index]

			if b_char_freqs > a_char_freqs {
				answer += b_char_freqs - a_char_freqs
			}
		}
		fmt.Fprintf(w, "%v\n", answer)
	}
}
