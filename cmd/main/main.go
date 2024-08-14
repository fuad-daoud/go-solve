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
	const max_n = 1000
	var n, k int
	fmt.Fscanf(reader, "%v %v\n", &n, &k)

	grid := [max_n]string{}

	for index := 0; index < n; index++ {
		fmt.Fscanf(reader, "%v\n", &grid[index])
	}
	answer := [max_n][max_n]rune{}
	answer_rows := 0
	for rows_index, row := range grid {
		if rows_index != 0 && rows_index%k != 0 {
			continue
		}
		answer_columns := 0
		for columns_index, column := range row {
			if columns_index != 0 && columns_index%k != 0 {
				continue
			}
			answer[answer_rows][answer_columns] = column
			answer_columns++
		}
		answer_rows++
	}
	for index := 0; index < n/k; index++ {
		for index2 := 0; index2 < n/k; index2++ {
			fmt.Fprintf(w, "%v", string(answer[index][index2]))
		}
		fmt.Fprintln(w)
	}
}
