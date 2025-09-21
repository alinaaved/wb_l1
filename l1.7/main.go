package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)

	var n int
	fmt.Scan(&n)

	var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := range n {
			mux.Lock()
			myMap[i] = n
			mux.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(myMap)
}
