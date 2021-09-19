package main

import "fmt"

func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	var localSize int
	for indexArr, arr := range matrix {
		for indexElem, elem := range arr {
			if elem == 1 {
				sizes = append(sizes, calculateThis(matrix, localSize, indexArr, indexElem))
			}
		}
	}
	return sizes
}

func calculateThis(matrix [][]int, localSize int, indexArr int, indexElem int) int {
	localSize++
	matrix[indexArr][indexElem] = 0
	if indexElem-1 >= 0 && matrix[indexArr][indexElem-1] == 1 {
		localSize = calculateThis(matrix, localSize, indexArr, indexElem-1)
	}
	if indexElem+1 <= len(matrix[indexArr])-1 && matrix[indexArr][indexElem+1] == 1 {
		localSize = calculateThis(matrix, localSize, indexArr, indexElem+1)
	}
	if indexArr-1 >= 0 && matrix[indexArr-1][indexElem] == 1 {
		localSize = calculateThis(matrix, localSize, indexArr-1, indexElem)
	}
	if indexArr+1 <= len(matrix)-1 && matrix[indexArr+1][indexElem] == 1 {
		localSize = calculateThis(matrix, localSize, indexArr+1, indexElem)
	}
	return localSize
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	sizes := RiverSizes(matrix)
	fmt.Println(sizes)
}
