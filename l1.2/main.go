package main

import (
	"fmt"
	"time"
)

func toSquare(num int) {
	fmt.Println(num * num)
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	for _, v := range array {
		go toSquare(v)
	}
	time.Sleep(time.Second) //время, чтобы горутины выполнились
}
