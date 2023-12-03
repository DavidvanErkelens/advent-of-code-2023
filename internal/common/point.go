package common

import "math"

type Point struct {
	X int
	Y int
}

type Point3D struct {
	X int
	Y int
	Z int
}

func (point *Point) ManhattanDistance(otherPoint Point) int {
	return Abs(point.X-otherPoint.X) + Abs(point.Y-otherPoint.Y)
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
