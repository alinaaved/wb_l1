package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {
	strLow := strings.ToLower(str)
	mapSigns := make(map[rune]bool)
	for _, v := range strLow {
		if mapSigns[v] == true {
			return false
		}
		mapSigns[v] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdeFAaf"
	str3 := "aabcd"

	res := checkUnique(str1)
	fmt.Println(res)
	res = checkUnique(str2)
	fmt.Println(res)
	res = checkUnique(str3)
	fmt.Println(res)

}
