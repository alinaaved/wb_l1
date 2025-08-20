package main

import "fmt"

type Human struct {
	field1 string
}

type Action struct {
	Human
	field2 string
}

func main() {
	var act Action = Action{Human{field1: "first"}, "second"}
	fmt.Println(act)
}
