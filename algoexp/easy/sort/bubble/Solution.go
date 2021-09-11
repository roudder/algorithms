// bubblesort
// Time comp.: 0(n2) space comp.: 0(1)

package main

import (
	"batman/src/algorithms/algoexp/easy/testhelper"
)
import "fmt"

func BubbleSort(array []int) []int {
	counter := 0
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-counter; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		counter += 1
	}
	return array
}
func main() {
	sorted := BubbleSort(testhelper.Array())
	fmt.Print(sorted)
}
