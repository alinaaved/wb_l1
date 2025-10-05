package main

import "fmt"

func main() {
	var arr1 []int
	var arr2 []int

	var num int

	fmt.Println("Введите элементы первого массива, 0 для выхода")
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		arr1 = append(arr1, num)
	}
	fmt.Println("Введите элементы второго массива, 0 для выхода")
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		arr2 = append(arr2, num)
	}

	mapNums := make(map[int]bool)
	for _, v := range arr1 {
		mapNums[v] = true
	}
	fmt.Println("Result: ")
	for _, v := range arr2 {
		if mapNums[v] {
			fmt.Print(v, " ")
		}
	}
}
