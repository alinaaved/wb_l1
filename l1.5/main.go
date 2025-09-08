package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	stop := time.After(10 * time.Second) //установлено время вручную - 10 сек
	wg.Add(1)
	go func() { //горутина пишет в канал
		defer wg.Done()
		var num int
		for {
			select {
			case <-stop:
				close(ch)
				return
			default:
				_, err := fmt.Scan(&num)
				if err != nil {
					close(ch)
					return
				}
				ch <- num
			}
		}
	}()
	wg.Add(1)
	go func() { //горутина читает канал
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
