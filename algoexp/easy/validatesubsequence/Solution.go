// sub is part of array?
// Time comp.: 0(n) space comp.: 0(1)

package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	index := 0
	for _, value := range array {
		isValid, ok := checkIndex(sequence, index)
		if ok {
			return isValid
		}
		if sequence[index] == value {
			index = index + 1
		}
	}
	if isValid, ok := checkIndex(sequence, index); ok {
		return isValid
	}
	return false
}

func checkIndex(sequence []int, index int) (bool, bool) {
	if len(sequence) == index {
		return true, true
	}
	return false, false
}

func main() {
	array := []int{1, 2, 3, 4}
	sequence := []int{0, 2, 3, 5}
	fmt.Print(IsValidSubsequence(array, sequence))
}
