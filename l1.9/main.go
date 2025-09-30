package main

import (
	"fmt"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
// во второй – результат операции x*2. После этого данные из второго канала должны
// выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами:
// генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно
// завершается.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range n {
		fmt.Scan(&arr[i])
	}

	go readArr(ch1, arr)
	go doubleArr(ch1, ch2)
	fmt.Println("Результат:")
	for j := range ch2 {
		fmt.Println(j)
	}
}

func readArr(ch1 chan int, arr []int) {
	for _, v := range arr {
		ch1 <- v
	}
	close(ch1)
}

func doubleArr(ch1 chan int, ch2 chan int) {
	for i := range ch1 {
		ch2 <- i * 2
	}
	close(ch2)
}
