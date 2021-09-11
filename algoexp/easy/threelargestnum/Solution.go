package main

import (
	"fmt"
	"math"
	"sort"
)

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	a := math.MinInt64
	b := math.MinInt64
	c := math.MinInt64

	arr := []int{a, b, c}

	for _, v := range array {
		justDoit(v, &arr)
	}
	return arr
}

func justDoit(v int, arr *[]int) {
	ar := *arr
	if v > ar[0] {
		ar[0] = v
		sort.Ints(ar)
	}
}
func main() {
	arr := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	numbers := FindThreeLargestNumbers(arr)
	fmt.Print(numbers)
}
