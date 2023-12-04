package main

import "fmt"

var justString string

/*
Если я правильно понял, то проблема в том, что если последний индекс строки меньше 100,
то мы ничего не получаеи на выход
Поэтому, чтобы вся строка выводилась, значением крайнего элемента, не включая этот элемент, будет длина строки
*/
func someFunc() {
	v := "somestring"
	a := len(v)
	justString = v[:a]
	fmt.Print(justString)
}

func main() {
	someFunc()
}
