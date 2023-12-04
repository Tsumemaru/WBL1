package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mas := [10]int{1, 5, 14, 7, 23, 1234, 6, 7, 24, 65}
	ch1 := make(chan int, 10)
	ch2 := make(chan int)
	// В цикле создаем буферизированный канал, иначе дедлок, так как значение из канала еще не читается
	for i := range mas {
		ch1 <- mas[i]
	}
	wg.Add(2)
	// Первая горутина, получающая значения
	go func() {
		var x int
		var i int = 0
		for {
			select {
			// Получаем значение из канала ch1 и отправляем в сh2 для другой горутины
			case x = <-ch1:
				ch2 <- x
				i++
			// Дефолт нужен, чтобы случайный кэйс не был обработан
			default:
			}
			/*
				Так как у нас массив, то завершение горутину можно сделать по счетчику(в нашем случае массив из 10 элементов
				и после того, как 10 раз прочитаем, завершаем горутину)
			*/
			if i == 10 {
				close(ch1)
				wg.Done()
				return
			}
		}
	}()
	// Вторая горутина, значение из ch2
	go func() {
		var i int = 0
		for {
			select {
			case y := <-ch2:
				fmt.Fprint(os.Stdout, y*y)
				i++
			default:
			}
			// Так же после 10 прочтений завершаем горутину
			if i == 10 {
				close(ch2)
				wg.Done()
				return
			}

		}

	}()
	wg.Wait()
}
