// # find some elements which are sum of target_sum value
// Time comp.: 0(n*log(n)) Memory comp.: 0(1)

package main

import (
	"fmt"
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	var left int = 0
	var right int = len(array) - 1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum < target {
			left = left + 1
		} else if currentSum > target {
			right = right - 1
		} else {
			return []int{array[left], array[right]}
		}
	}
	return []int{}
}
func main() {
	result := TwoNumberSum([]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5)
	if len(result) == 0 {
		fmt.Print("where is it?")
	} else {
		fmt.Print(result)
	}
}
