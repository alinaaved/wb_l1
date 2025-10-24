package main

import "fmt"

func reverse(str string) string {
	runeStr := []rune(str)
	for i := 0; i < (len(runeStr) / 2); i++ {
		runeStr[i], runeStr[len(runeStr)-i-1] = runeStr[len(runeStr)-i-1], runeStr[i]
	}
	return string(runeStr)
}

func main() {
	var str string
	fmt.Print("Enter string: ")
	fmt.Scanln(&str)
	reversedStr := reverse(str)
	fmt.Println("Result: ", reversedStr)
}
