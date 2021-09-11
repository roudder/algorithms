package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	if len(document) > len(characters) {
		return false
	}
	coincidences := make(map[string]int)
	for _, char := range document {
		str := fmt.Sprintf("%c", char)
		coincidences[str]++
	}
	for _, char := range characters {
		str := fmt.Sprintf("%c", char)
		if coincidences[str] > 0 {
			coincidences[str]--
			if coincidences[str] == 0 {
				delete(coincidences, str)
			}
		}
	}
	if len(coincidences) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	characters := "abd oby"
	doc := "bad boy"
	possible := GenerateDocument(characters, doc)
	fmt.Println(possible)
}
