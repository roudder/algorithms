package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array SpecialArray) int {
	sum := 0
	sum += ProductSumRec(array, 1)

	return sum
}

func ProductSumRec(i SpecialArray, depth int) int {
	sum := 0
	for _, value := range i {
		switch value.(type) {
		case int:
			sum += value.(int)
		default:
			if value, ok := value.(SpecialArray); ok {
				sum += ProductSumRec(value, depth+1)
			}
		}
	}
	return depth * sum
}

func main() {
	input := []interface{}{
		5, 2,
		[]int{7, -1},
		[]interface{}{7, []interface{}{7, -1}},
		3,
		SpecialArray{
			6,
			[]interface{}{
				-13, 8,
			},
			4,
		},
	}

	sum := ProductSum(input)
	fmt.Print(sum)

}
