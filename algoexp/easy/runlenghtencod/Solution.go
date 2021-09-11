package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func RunLengthEncoding(str string) string {
	var res bytes.Buffer
	var inner int
	var el string
	for i, v := range str {
		s := fmt.Sprintf("%c", v)
		if i == 0 {
			el = s
		}
		if el != s || inner == 9 {
			res.WriteString(strconv.Itoa(inner) + el)
			el = s
			inner = 1
			continue
		}
		inner++

	}
	res.WriteString(strconv.Itoa(inner) + el)
	return res.String()
}

func main() {
	encoding := RunLengthEncoding("AAAAAAAAAAAAAAABBBCCCDD")
	fmt.Print(encoding)
}
