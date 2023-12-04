package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		/*
			При работе с мапой, может возникнуть проблема, что при повторении ключа, значение перезаписывается
			поэтому важно,чтобы названия ключей были уникальны в целом, а не для одной горутины
		*/
		go func(counters map[int]int, i int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				/*
					Мьютекс лочится, когда какая-то горутина уже начала работу с мапой, только после завершения
					этой работы, получаем анлок для других горутин.
					Без мьютекса из-за состояния гонки получим фатал
				*/
				mu.Lock()
				counters[i*10+j] = j
				mu.Unlock()
			}
			wg.Done()
		}(counters, i, mu)
	}
	wg.Wait()
	fmt.Println("Результат", counters)
}
