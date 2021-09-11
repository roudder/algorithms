package main

import "fmt"

//could be done better without copy
//todo refactor it
func MoveElementToEnd(array []int, toMove int) []int {
	arrayCopy := []int{}
	count := 0

	for _, v := range array {
		if v != toMove {
			arrayCopy = append(arrayCopy, v)
		} else {
			count++
		}
	}
	for i := 0; i < count; i++ {
		arrayCopy = append(arrayCopy, toMove)
	}

	return arrayCopy
}

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	end := MoveElementToEnd(array, 2)
	fmt.Println(end)
}
