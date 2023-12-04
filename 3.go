import (
	"fmt"
	"sync"
)

// Передаем функции массив и индекс, получаем произведение элемента на себя
func square(mas [5]int, i int) int {
	b := mas[i] * mas[i]
	return b
}

func main() {
	var wg sync.WaitGroup
	mas := [5]int{2, 4, 6, 8, 10}
	/*
		Создаем буферизированный канал, чтобы избежать дэдлока, связанным с записью в канал,
		но если убрать waitgroup, то можно убрать буфер в канале, но основная горутина будет блокироваться,пока не прочитает всё из канала
	*/
	ch := make(chan int, 5)
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//Результат подсчета элемента массива отправляем в канал
			ch <- square(mas, i)
		}(i)
	}
	wg.Wait()
	//Получение суммы из канала
	answer := <-ch + <-ch + <-ch + <-ch + <-ch
	fmt.Println(answer)
}