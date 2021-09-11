package main

import "fmt"

// O(2^n) time | O(n) space
func getNthFib(n int) int {
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 0
	} else {
		return getNthFib(n-1) + getNthFib(n-2)
	}
}

// O(n) time | O(n) space
func getNthFib2(n int, m map[int]int) (int, map[int]int) {
	value, ok := m[n]
	if ok {
		return value, m
	}
	if n == 2 {
		return 1, m
	}
	if n == 1 {
		return 0, m
	} else {
		v1, _ := getNthFib2(n-1, m)
		v2, _ := getNthFib2(n-2, m)
		m[n] = v1 + v2
		return m[n], m
	}
}

// O(n) time | O(1) space
func getNthFib3(n int) int {
	if n == 0 {
		return 0
	}
	arr := []int{0, 1}
	counter := 3
	for counter <= n {
		next := arr[0] + arr[1]
		arr[0] = arr[1]
		arr[1] = next
		counter += 1
	}

	return arr[1]
}

func main() {
	fmt.Println(getNthFib(4))
	//todo think about transfer this part into method
	m := make(map[int]int)
	m[1] = 0
	m[2] = 1
	value, _ := getNthFib2(6, m)
	fmt.Println(value)
	fmt.Println(getNthFib3(0))
}
