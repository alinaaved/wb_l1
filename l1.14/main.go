package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	types(1)
	types("meow")
	types(true)
	types(ch)
}

func types(in interface{}) {
	switch i := in.(type) {
	case int:
		fmt.Println(i, " - число")
	case string:
		fmt.Println(i, " - строка")
	case bool:
		fmt.Println(i, " - булево значение")
	case chan int:
		fmt.Println(i, " - канал")
	}
}
