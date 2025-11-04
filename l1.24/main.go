package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными (приватными)
// полями x, y (типа float64) и конструктором. Расстояние рассчитывается по формуле
// между координатами двух точек.

//Подсказка: используйте функцию-конструктор NewPoint(x, y),
// Point и метод Distance(other Point) float64.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(p2 Point) float64 {
	x := p2.x - p.x
	y := p2.y - p.y

	return math.Sqrt(x*x + y*y)
}

func main() {
	var x1, x2, y1, y2 float64
	fmt.Print("Введите координаты первой точки: ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("Введите координаты второй точки: ")
	fmt.Scanln(&x2, &y2)

	point1 := NewPoint(x1, y1)
	point2 := NewPoint(x2, y2)

	res := point1.Distance(point2)
	fmt.Println("расстояние: ", res)
}
