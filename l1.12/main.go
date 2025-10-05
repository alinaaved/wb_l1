package main

import "fmt"

func main() {
	mapSet := make(map[string]bool)
	var word string
	fmt.Println("Введите строки, 0 для завершения")
	for {
		fmt.Scan(&word)
		if word == "0" {
			break
		}
		mapSet[word] = true
	}
	fmt.Print("Result:")
	for i := range mapSet {
		fmt.Print(" ", i)
	}
}
