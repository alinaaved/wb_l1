package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(start, end int, arr []rune) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		fmt.Println(start, end)
		start++
		end--
	}
}

func main() {
	fmt.Print("Введите строку: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	if len(str) > 0 && (str[len(str)-1] == '\n' || str[len(str)-1] == '\r') {
		str = str[:len(str)-1]
	}

	runeStr := []rune(str)
	reverse(0, len(runeStr)-1, runeStr)

	ind := 0

	for i := 0; i <= len(runeStr); i++ {
		if i == len(runeStr) || runeStr[i] == ' ' {
			reverse(ind, i-1, runeStr)
			ind = i + 1
		}
	}
	fmt.Println(string(runeStr))
}
