package main

import "fmt"

//todo add overlaping case
func BalancedBrackets(s string) bool {
	if len(s) == 0 {
		return false
	}
	m := make(map[string]int, 3)
	for _, rune := range s {
		switch string(rune) {
		case "{":
			m["{"]++
		case "}":
			m["{"]--
		case "[":
			m["["]++
		case "]":
			m["["]--
		case "(":
			m["("]++
		case ")":
			m["("]--
		}
	}
	if m["{"] == 0 && m["("] == 0 && m["["] == 0 {
		return true
	} else {
		return false
	}
	return false
}

func main() {
	s := "{}(){{}}"

	balanced := BalancedBrackets(s)
	fmt.Print(balanced)
}