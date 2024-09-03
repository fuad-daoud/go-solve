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

	template := make([]int, n)
	for index := 0; index < n; index++ {
		fmt.Fscanf(r, "%d", &template[index])
		if index != n-1 {
			fmt.Fscanf(r, " ")
		}
	}
	fmt.Fscanf(r, "\n")
	var m int
	fmt.Fscanf(r, "%d\n", &m)
	for index := 0; index < m; index++ {
		var x string
		fmt.Fscanf(r, "%s\n", &x)
		if len(x) != n {
			fmt.Fprintln(w, "NO")
			continue
		}
		yes := true
		mappings := make(map[byte]int)
		mappings2 := make(map[int]byte)
		for i := 0; i < n; i++ {
			{
				key, ok := mappings[x[i]]
				if !ok {
					mappings[x[i]] = template[i]
				} else if key != template[i] {
					fmt.Fprintln(w, "NO")
					yes = false
					break
				}
			}
			{
				newKey, ok := mappings2[template[i]]
				if !ok {
					mappings2[template[i]] = x[i]
				} else if newKey != x[i] {
					fmt.Fprintln(w, "NO")
					yes = false
					break
				}
			}

		}
		if yes {
			fmt.Fprintln(w, "YES")
		}
	}

}
