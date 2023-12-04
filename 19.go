package main

import "fmt"

func main() {
	var str string = "главрыба"
	/*
		Строку переводим в слайс рун, хотя можно и байт, зависит от символов(если английские буквы, то можно в байт,
		но русские буквы могут занимать больше одного байта, поэтому руны)
	*/
	strToRune := []rune(str)
	for i, j := 0, len(strToRune)-1; i < j; i, j = i+1, j-1 {
		// Просто меняем местами значение по крайним индексам, а потом переходим на другие
		// индексы, смещаясь к середине
		strToRune[i], strToRune[j] = strToRune[j], strToRune[i]
	}
	fmt.Print(string(strToRune))
}
