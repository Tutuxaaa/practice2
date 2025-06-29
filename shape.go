package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64 // Метод вычисления площади
	Name() string  // Метод для получения названия фигуры
}

// структура прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализация метода Area() для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Реализация метода Name() для Rectangle
func (r Rectangle) Name() string {
	return "Прямоугольник"
}

// структура круга
type Circle struct {
	Radius float64
}

// Реализация метода Area() для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Реализация метода Name() для Circle
func (c Circle) Name() string {
	return "Круг"
}

// функция для вывода площади любой фигуры
func PrintArea(s Shape) {
	fmt.Printf("%s: площадь = %.2f\n", s.Name(), s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	PrintArea(rect)
	PrintArea(circle)
}
