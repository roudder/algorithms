package main

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] == str[len(str)-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	str := "pppappp"
	println(IsPalindrome(str))
}
