package main

import "fmt"

//O()log()n times, O(1) space
func BinarySearch(array []int, target int) int {
	return BinarySearchRec(array, target, 0, len(array)-1)
}

func BinarySearchRec(array []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if target > array[middle] {
		start = middle - 1
	} else if target < array[middle] {
		end = middle + 1
	} else if target == array[middle] {
		return middle
	}
	res := BinarySearchRec(array, target, start, end)
	return res
}

func main() {
	output := BinarySearch([]int{0, 1, 21, 32, 33, 45, 45, 61, 71, 72, 73}, 33)
	fmt.Print(output)
}
