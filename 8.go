package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var i int64
	/*
		Получения числа в бинарном виде
	*/
	fmt.Print("Введите число\n")
	fmt.Fscan(os.Stdin, &i)
	// Тут получаем слайс нашего числа в виде 0 и 1
	x := intToBit(i)
	fmt.Println(x)
	fmt.Println("Укажите какой бит и какое значени вы хотите поместить")
	var a, b int64
	fmt.Fscan(os.Stdin, &a, &b)
	if b != 0 && b != 1 {
		fmt.Println("Указали неправильное значение бита")
		return
	}
	l := int64(len(x))
	if a >= l {
		fmt.Println("out of range")
		return
	}
	// Меняем значение элемента слайса
	x[a] = b
	fmt.Println(x)
	var newnumb int64 = 0
	// Переводим обратно в десятичную
	i = bitToInt(x, l, newnumb)
	fmt.Println(i)

}

// Функция перевода с десятичной в двоичную
func intToBit(i int64) []int64 {
	s := make([]int64, 0)
	for {
		if i%2 == 0 && i != 1 && i != 0 {
			s = append(s, i%2)
			i = i / 2
		} else if i%2 != 0 && i != 1 && i != 0 {
			s = append(s, i%2)
			i = (i - i%2) / 2
		} else if i == 1 {
			s = append(s, i)
			i = i - i
		} else {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}
	}
}

// Функция перевода с двоичной в десятичную
func bitToInt(x []int64, l int64, newnumb int64) int64 {
	for i := range x {
		newnumb += x[i] * int64(math.Pow(2, float64(l)-1))
		l--
	}
	return newnumb
}
