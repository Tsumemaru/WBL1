package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var gorNum int

type Counter struct {
	/*
		Не понял по условию задачи касательно того, как горутины должны работать со счетчиком
		Если не нужно запись защищать от других горутин, то просто убрать мьютекс
	*/
	mu    sync.Mutex
	count int
}

func main() {
	wg := sync.WaitGroup{}
	// Инициализируем структуру счетчик
	counter := Counter{count: 0}
	// Число горутин
	gorNum = 10
	// Через 10 секунд от инициализации контекста запишется в канал
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for i := 0; i < gorNum; i++ {
		go func(ctx context.Context, counter *Counter) {
			wg.Add(1)
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				// По дефолту будем инкрементировать счетчик, а как придет контекст, закончим записывать
				default:
					counter.mu.Lock()
					counter.count++
					counter.mu.Unlock()
				}
			}
		}(ctx, &counter)
	}
	wg.Wait()
	// Если без мьютексов, то значение получится раза примерно в 2,5 больше(с данными числом горутин)
	fmt.Print(counter.count)

}
