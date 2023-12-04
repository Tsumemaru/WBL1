package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b uint64
	// Получаем два числа
	fmt.Fscan(os.Stdin, &a, &b)
	// Если они меньше, то числа не соответствуют
	if a < uint64(math.Pow(2, 20)) || b < uint64(math.Pow(2, 20)) {
		fmt.Println("Ввели значения не соответствующие условию")
		return
	}
	/*
	 Если же ввели корректные числа, то также вводим операцию с этими числами
	 Uint64 имеют верхнуюю границу
	 И это касается не только чисел, которые вводятся, но и сам результат
	 если число и ожидаемый результат больше 18 446 744 073 709 551 615, то получим некорректный ответ
	*/
	var operation string
	_, err := fmt.Fscan(os.Stdin, &operation)
	if err != nil {
		fmt.Println("Ошибка ввода")
	}
	if operation == "*" || operation == "/" || operation == "+" || operation == "-" {
		switch operation {
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		}
	} else {
		fmt.Println("неверная операция")
	}
}
