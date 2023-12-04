package main

import (
	"fmt"
	"math"
	"os"
)

// Структура точек
type Point struct {
	X float64
	Y float64
}

// Функция подсчета дистанции между точками
func distance(A *Point, B *Point) float64 {
	dist := math.Sqrt(math.Pow((B.X-A.X), 2) + math.Pow((B.Y-A.Y), 2))
	return dist
}

// Конструктор, для которого вводим координаты точек через консоль
func NewPoint() *Point {
	p := new(Point)
	var a, b float64
	fmt.Fscan(os.Stdin, &a, &b)
	p.X = a
	p.Y = b
	return p
}

func main() {
	fmt.Print("Введите координаты первой точки\n")
	A := NewPoint()
	fmt.Print("Введите координаты второй точки\n")
	B := NewPoint()
	fmt.Println("Расстояние между точками: ", distance(A, B))
}
