package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	if len(str) == 1 {
		return 0
	}
	for i := range str {
		for j := range str {

			if j == len(str)-1 {
				if string(str[i]) != string(str[j]) || i == len(str)-1 {
					return i
				}
			}
			if i == j {
				continue
			}
			if string(str[j]) == string(str[i]) {
				break
			}
		}
	}
	return -1
}

func main() {
	index := FirstNonRepeatingCharacter("aac")
	fmt.Println(index)
}
