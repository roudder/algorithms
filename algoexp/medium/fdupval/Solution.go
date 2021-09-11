package main

func FirstDuplicateValue(array []int) int {
	m := make(map[int]struct{})
	for _, v := range array {
		_, ok := m[v]
		if ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}
	return -1
}
