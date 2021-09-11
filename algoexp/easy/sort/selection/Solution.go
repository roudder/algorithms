package main

import "fmt"

//0(n2) time | O(1) space)
func SelectionSort(array []int) []int {
	for i := range array {
		min := i
		for j := i; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}
	return array
}
func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	sort := SelectionSort(array)
	fmt.Print(sort)
}
