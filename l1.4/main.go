package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, ch <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена", id)
			return
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("worker ", id, ": ", msg)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	n := -1
	for n < 1 {
		fmt.Print("Enter number of workers: ")
		fmt.Scan(&n)
	}

	ch := make(chan string, 10)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, ch, i, &wg)
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				cancel()
				return
			}
			select {
			case <-ctx.Done():
				return
			case ch <- line:
			}
		}
	}()

	<-ctx.Done()
	close(ch)
	wg.Wait()

}
