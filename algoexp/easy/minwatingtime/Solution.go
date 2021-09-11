package main

import (
	"fmt"
	"sort"
)

//O(nLogn) time | O(1) space
func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	sum := 0
	localSum := 0
	for i := range queries {
		if i == 0 {
			continue
		}
		localSum = queries[i-1] + localSum
		sum += localSum
	}
	return sum
}

func main() {
	arr := []int{3, 2, 1, 2, 6}
	fmt.Println(MinimumWaitingTime(arr))
}
