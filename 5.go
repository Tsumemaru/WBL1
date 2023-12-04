package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// После указанного времени программа завершается
	var duration time.Duration = 20 * time.Second
	ctx := context.Background()
	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	ch := make(chan int)
	defer close(ch)
	go func() {
		var a int
		/*
			select при невыполнении всех условий выполняется рандомный, но в данном случае
			у нас либо данные из канала приходят, либо сработает контекст(и данные из канала ch уже прочитаны)
		*/
		for {
			select {
			case a = <-ch:
				fmt.Println(a)
			case <-ctx.Done():
				fmt.Println("Cчитывание закончилось")
				return

			}
		}
	}()
Loop:
	// Пишем значения в канал(из-за слипа и записи в канал за 20 секунд отработает 18 раз)
	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			break Loop
		case ch <- i:
		}
	}
}
