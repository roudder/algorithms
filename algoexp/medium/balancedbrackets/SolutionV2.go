package main

import "fmt"

var mapping = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
}

func BalancedBracketsV2(s string) bool {
	var arr []string
	for _, rune := range s {
		str := string(rune)
		switch str {
		case "{", "[", "(":
			arr = append(arr, str)
		case "}", ")", "]":
			if len(arr) > 0 && arr[len(arr)-1] == mapping[str] {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}

func main() {
	s2 := "([])(){}(())()())"
	balanced := BalancedBracketsV2(s2)
	fmt.Print(balanced)
}
