package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	array := []int{0, 1, 2, 2, 3, 0, 4, 2}
	length := removeElement(array, 2)
	fmt.Println(length)
	fmt.Println(array)

	// data := read()
}

func read() string {
	data, err := os.ReadFile("/home/fuad/projects/go-solve/cmd/main/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(data)
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	counter := 0
	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i] = math.MaxInt
			counter++
		}
	}
	slices.Sort(nums)
	return n - counter
}
