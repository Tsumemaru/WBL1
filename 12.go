package main

import "fmt"

func main() {
	a1 := []string{"cat", "cat", "dog", "cat", "tree"}
	// Переопределяем a1
	a1 = newSet(a1)
	fmt.Println(a1)
}

func newSet(a1 []string) []string {
	//Проверка по мапе с булевыми значениями
	keys := make(map[string]bool)
	a2 := []string{}
	for _, str := range a1 {
		/*
			Получаем булевое значение,
			Eсли нет такого ключа str, то получаем false и проходим по условию
			Если есть такой ключ str, то получаем true и не проходим по условию
		*/
		if value := keys[str]; !value {
			// Устанавливаем для этого ключа значение true,
			// Чтобы следующая  строка равная одной из предыдущих не проходила условие
			keys[str] = true
			// Пишем значение в слайс
			a2 = append(a2, str)
		}
	}
	return a2
}
