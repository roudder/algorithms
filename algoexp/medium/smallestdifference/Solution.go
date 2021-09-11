package main

import (
	"fmt"
	"math"
	"sort"
)

// O(nLog(n)) + mLog(m) time | O(1)
func SmallestDifference(array1, array2 []int) []int {
	smallest := math.MaxInt64
	current := math.MaxInt64
	smallestPair := []int{}
	i, j := 0, 0
	sort.Ints(array1)
	sort.Ints(array2)
	for i < len(array1) && j < len(array2) {
		i1, i2 := array1[i], array2[j]
		if i1 < i2 {
			i++
			current = i2 - i1
		} else if i1 > i2 {
			j++
			current = i1 - i2
		} else {
			return []int{i1, i2}
		}
		if current < smallest {
			smallest = current
			smallestPair = []int{i1, i2}
		}
	}
	return smallestPair
}

func main() {
	a1 := []int{10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123}
	a2 := []int{-1441, -124, -25, 1014, 1500, 660, 410, 245, 530}
	difference := SmallestDifference(a1, a2)
	fmt.Println(difference)
}
