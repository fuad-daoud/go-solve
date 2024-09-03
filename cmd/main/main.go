package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	solver := Solver{reader, writer}
	solver.Solve()
}

type Solver struct {
	reader io.Reader
	writer *bufio.Writer
}

func (solver Solver) Solve() {
	r := solver.reader
	w := solver.writer

	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Print("faced an error flushing " + err.Error())
			panic(err)
		}
	}(w)

	var cases int
	fmt.Fscanf(r, "%d\n", &cases)
	solver.pre_solve()
	for testCaseNumber := 0; testCaseNumber < cases; testCaseNumber++ {
		solver.solveCase(testCaseNumber)
	}

}

func (solver Solver) pre_solve() {
}

func (solver Solver) solveCase(case_number int) {

	r := solver.reader
	w := solver.writer

	var n, k int

	fmt.Fscanf(r, "%d %d\n", &n, &k)

	numbers := make([]struct {
		first  int64
		second int64
	}, n)

	for index := 0; index < int(n); index++ {
		fmt.Fscanf(r, "%d", &numbers[index].first)
	}
	fmt.Fscanf(r, "\n")
	for index := 0; index < int(n); index++ {
		fmt.Fscanf(r, "%d", &numbers[index].second)
	}
	fmt.Fscanf(r, "\n")
	sort.Slice(numbers, func(a, b int) bool {
		return numbers[a].first < numbers[b].first
	})
	var answer int64 = 0

	for index := 0; index < int(n); index++ {
		if numbers[index].second == 1 {
			var median int64

			if index < int(n)/2 {
				median = numbers[n/2].first
			} else {
				median = numbers[(n-2)/2].first
			}
			var current_answer int64 = median + int64(k) + numbers[index].first

			if current_answer > answer {
				answer = current_answer
			}
		}
	}

	var left, right int64 = 0, 2_000_000_000

	for {
		if left >= right {
			break
		}
		var mid int64 = (1 + right + left) / 2
		numbers_more_than_or_equal_mid := 0
		smaller_numbers := make([]int, 0)
		for index := 0; index < int(n)-1; index++ {
			if numbers[index].first >= mid {
				numbers_more_than_or_equal_mid++
			} else if numbers[index].second == 1 {
				smaller_numbers = append(smaller_numbers, int(mid-numbers[index].first))
			}
		}
		sort.Slice(smaller_numbers, func(i, j int) bool {
			return smaller_numbers[i] < smaller_numbers[j]
		})
		fmt.Printf("smaller_numbers: %v\n", smaller_numbers)
		temp_k := k
		for _, value := range smaller_numbers {
			if temp_k >= value {
				temp_k -= value
				numbers_more_than_or_equal_mid++
			}
		}
		if numbers_more_than_or_equal_mid >= (n+1)/2 {
			left = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Printf("left: %v\n", left)
	fmt.Printf("numbers: %v\n", numbers)
	current_answer := numbers[n-1].first + left

	if current_answer > answer {
		answer = current_answer
	}

	fmt.Fprintf(w, "%d\n", answer)
}
