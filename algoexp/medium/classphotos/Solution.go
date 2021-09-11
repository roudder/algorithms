package main

import (
	"fmt"
	"sort"
)

// ClassPhotos O(nlog(n)) time | O(1) space
func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	var redUp bool
	if redShirtHeights[0] > blueShirtHeights[0] {
		redUp = false
	} else if redShirtHeights[0] < blueShirtHeights[0] {
		redUp = true
	} else {
		return false
	}
	for i := range redShirtHeights {
		if redShirtHeights[i] < blueShirtHeights[i] && redUp == true {
			continue
		} else if redShirtHeights[i] > blueShirtHeights[i] && redUp == false {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	red := []int{5, 8, 1, 3, 4}
	blue := []int{6, 9, 2, 4, 5}
	fmt.Printf("%v", ClassPhotos(red, blue))
}
