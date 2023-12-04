package main

import (
	"fmt"
	"os"
	"strings"
)

func check(str string) bool {
	// Сначала переводим все в нижний регистр
	str = strings.ToLower(str)
	// Потом записываем в слайс рун
	newrune := []rune(str)
	// Циклы в которых сравниваем элементы слайса, пока не найдет одинаковые
	// Если находит то false
	for i := range newrune {
		for j := range newrune {
			if i != j && newrune[i] == newrune[j] {
				return false
			}
		}
	}
	// Если нет то true
	return true
}

func main() {
	var str string
	// Ввести строку через консоль(конечно, чтобы там были символы, которые могут быть рунами)
	fmt.Fscan(os.Stdin, &str)
	fmt.Print(check(str))
}
