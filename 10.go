package main

import (
	"fmt"
	"math"
)

func main() {
	mas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// i является шагом
	for i := -20.0; i < 40; {
		fmt.Println(i, " : ", calc(mas, i))
		i = i + 10
	}
}

// Из функции получаем слайс, который соответствует шагу
func calc(mas []float64, i float64) []float64 {
	a := make([]float64, 0, 3)
	for y := range mas {
		/*
			 Не получать при math.Mod результаты которые могут отличаться от самого шага на +10 и больше
			Проверка, чтобы не было случаев, когда значение положительное, а шаг отрицательный и наоборот
		*/
		if math.Mod(mas[y], i) < 10 && mas[y]/i > 0 && mas[y]-i < 10 && mas[y]/i < 2 {
			a = append(a, mas[y])
		}

	}
	return a
}
