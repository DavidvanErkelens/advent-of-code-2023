package aoc_grid

import (
	"advent-of-code-2023/internal/helpers"
	"fmt"
)

type Grid[T comparable] struct {
	Width      int
	Height     int
	values     [][]GridPoint[T]
	dataPoints [][]T
}

type GridPoint[T comparable] struct {
	X     int
	Y     int
	Value T
}

func NewRuneGrid(input string) Grid[rune] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	values := make([][]GridPoint[rune], 0)
	dataPoints := make([][]rune, 0)

	for y, line := range lines {
		runes := []rune(line)
		points := make([]GridPoint[rune], 0)
		for x, r := range runes {
			points = append(points, GridPoint[rune]{
				X:     x,
				Y:     y,
				Value: r,
			})
		}
		values = append(values, points)
		dataPoints = append(dataPoints, runes)
	}

	return Grid[rune]{
		Width:      width,
		Height:     height,
		values:     values,
		dataPoints: dataPoints,
	}
}

func NewRuneGridFromParsedData(input [][]rune) Grid[rune] {
	width := len(input[0])
	height := len(input)

	values := make([][]GridPoint[rune], 0)

	for y, line := range input {
		points := make([]GridPoint[rune], 0)
		for x, r := range line {
			points = append(points, GridPoint[rune]{
				X:     x,
				Y:     y,
				Value: r,
			})
		}
		values = append(values, points)
	}

	return Grid[rune]{
		Width:      width,
		Height:     height,
		values:     values,
		dataPoints: input,
	}
}

func NewIntGrid(input string, separator string) Grid[int] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	values := make([][]GridPoint[int], 0)
	dataPoints := make([][]int, 0)

	for y, line := range lines {
		numbers := helpers.StringListOfNumericValuesToSlice(line, separator)
		points := make([]GridPoint[int], 0)
		for x, number := range numbers {
			points = append(points, GridPoint[int]{
				X:     x,
				Y:     y,
				Value: number,
			})
		}
		values = append(values, points)
		dataPoints = append(dataPoints, numbers)
	}

	return Grid[int]{
		Width:      width,
		Height:     height,
		values:     values,
		dataPoints: dataPoints,
	}
}

func (g *Grid[T]) Lines() [][]T {
	return g.dataPoints
}

func (g *Grid[T]) SetValueAt(x, y int, val T) {
	g.values[y][x] = GridPoint[T]{
		X:     x,
		Y:     y,
		Value: val,
	}

	g.dataPoints[y][x] = val
}

func (g *Grid[T]) ValueAt(x int, y int) T {
	return g.values[y][x].Value
}

func (g *Grid[T]) ElementAt(x int, y int) GridPoint[T] {
	return g.values[y][x]
}

func (g *Grid[T]) InRange(x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= g.Width {
		return false
	}

	if y >= g.Height {
		return false
	}

	return true
}

func (g *Grid[T]) FindLocationOfValue(val T) (int, int) {
	for y, line := range g.values {
		for x, value := range line {
			if value.Value == val {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g *Grid[T]) GetNeighbors(x, y int) []GridPoint[T] {
	neighbors := make([]GridPoint[T], 0)
	toCheck := []struct{ x, y int }{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	for _, p := range toCheck {
		if g.InRange(p.x, p.y) {
			neighbors = append(neighbors, g.values[p.y][p.x])
		}
	}

	return neighbors
}

func (g *Grid[T]) Print(stringify func(T) string) {
	for _, line := range g.Lines() {
		for _, value := range line {
			fmt.Print(stringify(value))
		}
		fmt.Print("\n")
	}
}
