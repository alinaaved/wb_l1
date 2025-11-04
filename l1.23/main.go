package main

import (
	"fmt"
)

//Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

//Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
//(copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Введите номер элемента слайса для удаления: ")

	var i int
	fmt.Scanln(&i)
	if i < 0 || i > len(slice) {
		panic("Неверный индекс")
	}
	copy(slice[i-1:], slice[i:])
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
