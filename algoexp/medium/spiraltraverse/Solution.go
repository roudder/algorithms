package main

import "fmt"

// O (n) time | O(n) space
func SpiralTraverse(array [][]int) []int {
	result := []int{}
	startRow := 0
	endRow := len(array) - 1
	startColumn := 0
	endColumn := len(array[0]) - 1
	SpiralTraverseRec(startRow, endRow, startColumn, endColumn, &array, &result)
	return result
}

func SpiralTraverseRec(startRow int, endRow int, startColumn int, endColumn int, array *[][]int, result *[]int) {
	if startRow > endRow || startColumn > endColumn {
		return
	}
	for i := startColumn; i <= endColumn; i++ {
		*result = append(*result, (*array)[startRow][i])
	}

	for i := startRow + 1; i <= endRow; i++ {
		*result = append(*result, (*array)[i][endColumn])
	}

	if startRow != endRow {
		for i := endColumn - 1; i >= startColumn; i-- {
			*result = append(*result, (*array)[endRow][i])
		}
	}

	if startColumn != endColumn {
		for i := endRow - 1; i > startRow; i-- {
			*result = append(*result, (*array)[i][startColumn])
		}
	}

	SpiralTraverseRec(startRow+1, endRow-1, startColumn+1, endColumn-1, array, result)
}

func main() {
	array := [][]int{{1, 2, 3}, {12, 13, 4}, {11, 14, 5}, {10, 15, 6}, {9, 8, 7}}
	result := SpiralTraverse(array)
	fmt.Println(result)
}
