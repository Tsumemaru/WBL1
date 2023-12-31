package main

import (
	"fmt"
)

func main() {
	a := []int{23, 52, 154, 1, 65, 76, 89, 34, 22}
	fmt.Print(quicksort(a))
}

func quicksort(a []int) []int {
	/*
		Если в слайсе меньше двух элементов, то ничего не делаем
		Так же с помощью этой проверки останавливаем рекурсию
	*/
	if len(a) < 2 {
		return a
	}
	var n int
	// l и r крайние элементы слайса
	l, r := 0, len(a)-1
	// Выбираем индекс элемента в середине слайса(но в целом это необязательно)
	if (l+r)%2 == 0 {
		n = (l + r) / 2
	} else {
		n = (l + r + 1) / 2
	}
	// меняем местами серединный и правый элемент
	a[n], a[r] = a[r], a[n]
	// Каждый элемент сравнивается с a[r](это уже 65, если по этому примеру)
	for i := range a {
		/*
			Если элемент меньше a[r], то левый и выбранный элемент меняются местами и левым становится следующий элемент по индексу
			Таким образом мы сравниваем все элементы с правым и меняем местами, при рекурсивном вызове меняем крайний правый элемент каждый раз
		*/
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l]
			l++
		}
	}
	a[l], a[r] = a[r], a[l]
	//рекурсивно выполняем сортировку для кусков
	quicksort(a[:l])
	quicksort(a[l+1:])

	return a
}
