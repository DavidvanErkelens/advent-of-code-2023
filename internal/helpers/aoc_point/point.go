package aoc_point

import (
	math2 "advent-of-code-2023/internal/helpers/aoc_math"
	"math"
)

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Point struct {
	X int
	Y int
}

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{X: x, Y: y, Z: z}
}

type Point3D struct {
	X int
	Y int
	Z int
}

func (point *Point) ManhattanDistance(otherPoint Point) int {
	return math2.Abs(point.X-otherPoint.X) + math2.Abs(point.Y-otherPoint.Y)
}

func (point *Point) GradientTo(otherPoint Point) float64 {
	return float64(otherPoint.Y-point.Y) / float64(otherPoint.X-point.X)
}

func (point *Point) AngleTo(otherPoint Point) float64 {
	res := math.Atan2(float64(otherPoint.Y-point.Y), float64(otherPoint.X-point.X)) * 180 / math.Pi
	if res < 0 {
		res += 360
	}
	return res
}
