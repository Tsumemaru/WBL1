package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "snow dog sun"
	// С пакетом strings создаем слайс строк, состоящих из отдельных слов без пробелов
	s := strings.Split(str, " ")
	// Меняем местами так же, как и просто символы, местами
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	// Обратная операция Split, делаем строку из слайса с добавлением пробелов
	fmt.Println(strings.Join(s, " "))
}
