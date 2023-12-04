package main

import "fmt"

// Структура Human со встроенной структурой Action
type Human struct {
	Name string
	Age  string
	Action
}

// Структура часть структуры Human
type Action struct {
	Status string
}

// Метод, который отнсится к Action
func (a *Action) showStatus() {
	fmt.Print("На данный момент " + a.Status + " !")
}
func main() {
	// Создаем структуру Human
	human := new(Human)
	// Поле Status поле встроенной структуры Action(необязательно явно указывать структуру)
	human.Status = "жду"
	//Тоже относится к методам структуры Action, можем вызывать метод при обьявлении Human
	human.showStatus()
}
