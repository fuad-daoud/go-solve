package main

import (
	"fmt"
	"os"
)

func main() {
	array := []int{1, 2, 3, 4, 4, 3, 2, 1}
	newArray := shuffle(array, 4)
	fmt.Println(newArray)

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

func shuffle(nums []int, n int) []int {
	newArray := make([]int, 0)
	for i := 0; i < n; i++ {
		newArray = append(newArray, nums[i], nums[i+n])
	}
	return newArray
}
