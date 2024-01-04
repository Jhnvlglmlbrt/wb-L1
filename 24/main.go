package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку в пространстве
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p, q Point) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func main() {
	point1 := NewPoint(10.1, 15.2)
	point2 := NewPoint(-15.2, 3.1)

	distance := Distance(point1, point2)

	fmt.Printf("Расстояние между точкой (%.2f, %.2f) и точкой (%.2f, %.2f) равно %.2f\n", point1.x, point1.y, point2.x, point2.y, distance)
}
