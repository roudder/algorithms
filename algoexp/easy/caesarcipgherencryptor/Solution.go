package main

import (
	"strconv"
	"strings"
)

func caesarCipherEncryptor(str string, key int) string {
	var strArr []string
	key = key % 26
	for _, value := range str {
		ascinum := int(value) + key
		if ascinum > 122 {
			ascinum = ascinum - 26
		}
		strArr = append(strArr, strconv.Itoa(ascinum))
		if strArr == nil {
			panic("do smth, please")
		}
	}
	return strings.Join(strArr, "")
}

func main() {
	caesarCipherEncryptor("xyz", 1)
}
