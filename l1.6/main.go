package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//1. выход по условию
	stopAt := 3
	go func() {
		defer fmt.Println("отложенное выполнение до выхода")
		for i := 1; i <= 10; i++ {
			if i == stopAt {
				fmt.Println("условие для выхода выполнено")
				return // мягкая остановка
			}
			fmt.Println("горутина:", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second) //время для доработки

	//2. через канал уведомления
	done := make(chan bool)

	go func() {
		defer fmt.Println("остановка")
		for {
			select {
			case <-done:
				return //если есть сигнал остановки
			default:
				fmt.Println("в работе")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(750 * time.Millisecond)
	close(done) //останавливаем
	time.Sleep(200 * time.Millisecond)

	//3. через контекст(WithCancel)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		defer fmt.Println("ctx cancel -> return")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("работа")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(700 * time.Millisecond)
	cancel() //остановка
	time.Sleep(200 * time.Millisecond)

	//4. через контекст(WithTimeout)
	ctx2, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel() //если main выйдет раньше

	go func(ctx2 context.Context) {
		defer fmt.Println("timeout -> return")
		for {
			select {
			case <-ctx2.Done():
				return
			default:
				fmt.Println("работа")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx2)

	time.Sleep(time.Second)

	//5.runtime.Goexit()
	go func() {
		defer fmt.Println("отложенное выполнение перед выходом")
		fmt.Println("до выхода")
		runtime.Goexit()
		//сюда уже не попадаем
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("запущена горутина ", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
}
