package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// WaitGroup позволяет не заканчивать работу программы, пока работают другие горутины
	var wg sync.WaitGroup
	mas := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		//Каждый раз добавляем счетчику ожидания 1 при запуске новой горутины
		wg.Add(1)
		//Создаем горутину, передаем элемент массива, каждая горутина получит свой элемент
		go func(i int) {
			//Умножает этот элемент на себя
			fmt.Fprintln(os.Stdout, mas[i]*mas[i])
			//Убавляем значение счетчика на 1
			wg.Done()
		}(i)
	}
	// Как только счетчик будет равен нулю, блокировка спадет и основная горутина продолжит работу
	wg.Wait()
	fmt.Fprintln(os.Stdout, "Программа отработала")
}