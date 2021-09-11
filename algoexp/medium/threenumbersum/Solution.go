package main

import (
	"batman/src/algorithms/algoexp/easy/testhelper"
	"fmt"
	"sort"
)

//O(n^2) time | O(n) space
func ThreeNumberSum(array []int, target int) [][]int {
	finalArray := [][]int{}
	sort.Ints(array)
	for i := 0; i < len(array)-2; i++ {
		start := i
		left := i + 1
		right := len(array) - 1
		for left != right && left < right {
			sum := array[start] + array[left] + array[right]
			if sum < target {
				left += 1
			} else if sum > target {
				right -= 1
			} else {
				arr := []int{array[start], array[left], array[right]}
				finalArray = append(finalArray, arr)
				left += 1
				right -= 1
			}
		}
	}
	return finalArray
}

func main() {
	array := testhelper.Array1()
	final := ThreeNumberSum(array, 0)
	fmt.Println(final)
}
