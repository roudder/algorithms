package main

import (
	"sort"
)

// NonConstructibleChange O(nlogn) time | 0(1) space
func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	currentLine := 0
	for _, coin := range coins {
		if coin > currentLine+1 {
			return currentLine + 1
		}
		currentLine += coin
	}
	return currentLine + 1
}

func main() {
	arr := []int{5, 7, 1, 1, 2, 3, 22}
	print(NonConstructibleChange(arr))
}
