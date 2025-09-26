package main

import "fmt"

func main() {
	var n int64
	var b, val int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)
	fmt.Print("Введите номер бита: ")
	fmt.Scanln(&b)
	fmt.Print("Введите значение (0 или 1): ")
	fmt.Scanln(&val)

	if val == 1 {
		n |= 1 << (b - 1)
	} else {
		n &^= 1 << (b - 1)
	}

	fmt.Printf("%b", n)
}
