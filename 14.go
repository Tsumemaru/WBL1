package main

import "fmt"

func main() {
	// В интерфейс а можно поставить все по условию задачи
	var a interface{} = make(chan int)
	// Switch со v, который означает тип в интерфейсе
	// Единственное канал должен быть конкретного типа в кейсах и в интерфейсе
	switch v := a.(type) {
	case int:
		fmt.Print("Это инт ", v)

	case string:
		fmt.Print("Это строка ", v)

	case bool:
		fmt.Print("Это булево значение ", v)

	case chan int:
		fmt.Print("Это канал ", v)
	}
}
