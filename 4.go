package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var workersNumber int

func main() {
	var wg sync.WaitGroup
	// Создаем контекст, который дает сигнал через канал, если было нажато сочетание ctrl-c
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	// Число горутин, которое можно менять
	workersNumber = 10
	ch := make(chan int)
	defer close(ch)
	exitChan := make(chan struct{}, 1)
	defer close(exitChan)
	go func() {
		for i := 0; workersNumber > i; i++ {
			wg.Add(1)
			go func(i int) {
			LOOP:
				for {
					select {
					// Кейс по ожиданию сигнала от контескта
					case <-ctx.Done():
						wg.Done()
						break LOOP
					case a := <-ch:
						/*
							Из-за того, что stdout - поток вывода, не все горутины в состоянии конкурентости сразу смогут вывести свое значение,
							Только по завершению горутин остальных горутин выведут данные в поток и завершатся сами
						*/
						fmt.Fprintln(os.Stdout, a)
					}
				}
				fmt.Println("Горутина", i, "отработала")
				// Когда все горутины получат сигнал от контекста, то все закроют свои циклы
			}(i)
		}
	}()
	// Цикл, который работает пока нет контекста
LOOP1:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break LOOP1
		case ch <- i:
		}
	}
	//Как все горутины отработают, продолжим основную горутину
	wg.Wait()
	fmt.Println("Программа отработала")
}
