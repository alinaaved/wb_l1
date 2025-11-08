package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	if duration <= 0 {
		return
	}
	<-time.After(duration)
}

func main() {
	go func() {
		fmt.Println("горутина 1")
	}()
	sleep(1 * time.Second)
	fmt.Println("горутина мейн")
}
