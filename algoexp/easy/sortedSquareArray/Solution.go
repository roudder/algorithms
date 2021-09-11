package main

import (
	"fmt"
	"sort"
)

func SortedSquaredArray(array []int) []int {
	for i := range array {
		array[i] = array[i] * array[i]
	}
	sort.Ints(array)
	return array
}

func main() {
	ints := []int{-4, -2}
	ints = SortedSquaredArray(ints)
	fmt.Println(ints)
}
