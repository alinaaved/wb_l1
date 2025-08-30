package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range ch {
		fmt.Println("worker ", id, ": ", j)
		time.Sleep(time.Second)
	}
}

func main() {
	n := -1
	var data string
	for n < 1 {
		fmt.Print("Enter number of workers: ")
		fmt.Scan(&n)
	}

	ch := make(chan string, 10)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ch, i, &wg)
	}
	for {
		fmt.Scan(&data)
		if data == "exit" {
			break
		}
		ch <- data
	}
	close(ch)
	wg.Wait()

}
