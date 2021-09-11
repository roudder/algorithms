package testhelper

var array []int

func Array() []int {
	if array == nil {
		array = []int{-7, -5, -3, -1, 0, 1, 3, 5, 7}
	}
	return array
}

func Array1() []int {
	if array == nil {
		array = []int{12, 3, 1, 2, -6, 5, -8, 6}
	}
	return array
}
