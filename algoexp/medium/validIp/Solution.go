package main

import (
	"fmt"
	"strconv"
)

// ValidIPAddresses todo finalize "zero" cases like 3700100
func ValidIPAddresses(str string) []string {
	if len(str) < 4 || len(str) > 12 {
		return []string{}
	}
	res := make([]string, 0)
	someMagic(str, &res, "", 4)
	return res
}

func someMagic(str string, res *[]string, localStart string, chunk int) {
	if chunk == 0 && len(str) > 0 {
		return
	} else if chunk == 0 {
		intVal, _ := strconv.Atoi(str[:])
		if intVal <= 255 {
			*res = append(*res, localStart)
		}
		return
	}

	if len(str) == 0 {
		return
	}

	correctStart1 := str[0:1]
	var local1 string
	if localStart == "" {
		local1 = correctStart1
	} else {
		local1 = localStart + "." + correctStart1
	}
	p1 := chunk - 1
	someMagic(str[1:], res, local1, p1)
	if len(str) >= 2 {
		correctStart2 := str[0:2]
		var local2 string
		if localStart == "" {
			local2 = correctStart2
		} else {
			local2 = localStart + "." + correctStart2
		}
		p2 := chunk - 1
		someMagic(str[2:], res, local2, p2)
	}
	if len(str) >= 3 {
		intVal, _ := strconv.Atoi(str[0:3])
		if intVal <= 255 {
			correctStart3 := str[0:3]
			var local3 string
			if localStart == "" {
				local3 = correctStart3
			} else {
				local3 = localStart + "." + correctStart3
			}
			p3 := chunk - 1
			someMagic(str[3:], res, local3, p3)
		}
	}

}

func main() {
	str := "1921680"
	addresses := ValidIPAddresses(str)
	for _, address := range addresses {
		fmt.Println(address)
	}
	fmt.Println(len(addresses))
}
