// insertion sort
// Time comp.: 0(n2) space comp.: 0(1)

package main

import (
	"batman/src/algorithms/algoexp/easy/testhelper"
)
import "fmt"

func insertionSort(array []int) []int {
	for index := range array {
		inner := index
		for j := 0; j < inner; j++ {
			if array[index] < array[j] {
				array[index], array[j] = array[j], array[index]
				continue
			}
		}
	}
	return array
}
func main() {
	sorted := insertionSort(testhelper.Array())
	fmt.Print(sorted)
}
