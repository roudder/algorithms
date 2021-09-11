package main

import "fmt"

func IsMonotonic(array []int) bool {
	increase := true
	initialed := false

	if len(array) < 2 {
		return true
	}
	for i := 0; i < len(array)-1; i++ {
		if !initialed && array[i] < array[i+1] {
			increase = true
			initialed = true
		} else if !initialed && array[i] > array[i+1] {
			increase = false
			initialed = true
		}
		if (array[i] > array[i+1] && increase) || (array[i] < array[i+1] && !increase) {
			return false
		}
	}

	return true
}

func main() {
	array := []int{1, 1, 2, 3, 4, 3, 5, 5, 6, 7, 8, 8, 9, 10, 11}
	isMonotonic := IsMonotonic(array)
	fmt.Println(isMonotonic)
}
