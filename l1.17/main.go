package main

import "fmt"

func binSearch(arr []int, l, r int, el int) int {
	if l <= r {
		if arr[(l+r)/2] > el {
			r = (r+l)/2 - 1
			return binSearch(arr, l, r, el)
		} else if el == arr[(l+r)/2] {
			return (l + r) / 2
		} else if el > arr[(l+r)/2] {
			l = (r+l)/2 + 1
			return binSearch(arr, l, r, el)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println(binSearch(arr, 0, len(arr)-1, 2))
	fmt.Println(binSearch(arr, 0, len(arr)-1, 3))
	fmt.Println(binSearch(arr, 0, len(arr)-1, 9))
	fmt.Println(binSearch(arr, 0, len(arr)-1, 10))
}
