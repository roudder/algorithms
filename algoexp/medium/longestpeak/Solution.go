package main

import "fmt"

func LongestPeak(array []int) int {

	//O(n) time | space O(n)
	max := 0
	for i := 1; i < len(array)-1; i++ {
		isPeak := array[i] > array[i-1] && array[i] > array[i+1]
		if isPeak {
			left := i - 1
			right := i + 1
			for left > 0 && array[left-1] < array[left] {
				left -= 1
			}
			for right < len(array)-1 && array[right] > array[right+1] {
				right += 1
			}
			localMax := right - left + 1
			if localMax > max {
				max = localMax
			}
		}
	}
	return max
}
func main() {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	peak := LongestPeak(arr)
	fmt.Println(peak)
}
