package main

import (
	"fmt"
	"sync"
)

type increment struct {
	count int
}

func incrementing(m *sync.Mutex, c *increment, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	c.count++
}

func main() {
	var mux sync.Mutex
	incr := increment{count: 0}
	var n int
	fmt.Print("Введите количество итераций: ")
	fmt.Scanln(&n)
	var wg sync.WaitGroup

	for range n {
		wg.Add(1)
		go incrementing(&mux, &incr, &wg)
	}
	wg.Wait()
	fmt.Print(incr.count)
}
