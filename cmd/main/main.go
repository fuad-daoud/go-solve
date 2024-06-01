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
	var n, k, q int

	fmt.Fscanf(reader, "%d %d %d\n", &n, &k, &q)
	points, times, queries, answers := make([]int64, k), make([]int64, k), make([]Pair, q), make([]int64, q)

	for i := 0; i < k; i++ {
		if i == k-1 {
			fmt.Fscanf(reader, "%v\n", &points[i])
			break
		}
		fmt.Fscanf(reader, "%v ", &points[i])
	}
	for i := 0; i < k; i++ {
		if i == k-1 {
			fmt.Fscanf(reader, "%v\n", &times[i])
			break
		}
		fmt.Fscanf(reader, "%v ", &times[i])
	}
	for i := 0; i < q; i++ {
		fmt.Fscanf(reader, "%v\n", &queries[i].First)
		queries[i].Second = i
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i].First < queries[j].First
	})
	currentIndex := 0
	for _, query := range queries {
		for {
			if points[currentIndex] >= query.First {
				break
			}
			currentIndex++
		}
		previousPoint := getOrDefault(points, currentIndex-1)
		previousTime := getOrDefault(times, currentIndex-1)
		distance := points[currentIndex] - previousPoint
		time := times[currentIndex] - previousTime
		queryDistance := query.First - previousPoint
		queryTime := previousTime + queryDistance*time/distance
		answers[query.Second] = queryTime
	}
	for index, answer := range answers {
		if index == len(answers)-1 {
			fmt.Fprintf(w, "%d\n", answer)
			break
		}
		fmt.Fprintf(w, "%d ", answer)
	}
}

func getOrDefault(slice []int64, index int) int64 {
	if index >= len(slice) || index == -1 {
		return 0
	}
	return slice[index]
}

type Pair struct {
	First  int64
	Second int
}
