package main

import "fmt"

func partition(l, r int, a []int) int {
	mid := (r + l) / 2
	pivot := a[mid]
	a[mid], a[r-1] = a[r-1], a[mid] //временное перемещение пивота в конец
	i := l
	for j := l; j < r-1; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r-1] = a[r-1], a[i] //возвращение пивота
	return i
}

func sort(l, r int, a []int) {
	if r-l <= 1 {
		return
	}
	m := partition(l, r, a)
	sort(l, m, a)
	sort(m, r, a)
}

func main() {
	var elem int
	fmt.Print("Введите количество элементов: ")
	fmt.Scan(&elem)
	arr := make([]int, elem)
	fmt.Println("Введите элементы: ")
	for i := range elem {
		fmt.Scan(&arr[i])
	}
	sort(0, elem, arr)
	fmt.Println(arr)
}
