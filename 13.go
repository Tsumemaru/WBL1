package main

import (
	"fmt"
	"os"
)

func main() {
	list := [10]int{2, 5, 6, 34, 6, 1, 78, 1, 456, 4}
	fmt.Println("Укажите индекс двух чисел, которых хотите поменять местами")
	var a, b int
	fmt.Fscan(os.Stdin, &a, &b)
	// Так можно просто поменять значения местами, просто дать им другое значение
	list[a], list[b] = list[b], list[a]
	fmt.Println(list)
}
