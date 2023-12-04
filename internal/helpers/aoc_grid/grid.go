package aoc_grid

import "advent-of-code-2023/internal/helpers"

type Grid[T any] struct {
	Width  int
	Height int
	values [][]T
}

func NewRuneGrid(input string) Grid[rune] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	values := make([][]rune, 0)

	for _, line := range lines {
		values = append(values, []rune(line))
	}

	return Grid[rune]{
		Width:  width,
		Height: height,
		values: values,
	}
}

func NewIntGrid(input string, separator string) Grid[int] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	values := make([][]int, 0)

	for _, line := range lines {
		values = append(values, helpers.StringListOfNumericValuesToSlice(line, separator))
	}

	return Grid[int]{
		Width:  width,
		Height: height,
		values: values,
	}
}

func (g *Grid[T]) Lines() [][]T {
	return g.values
}

func (g *Grid[T]) ValueAt(x int, y int) T {
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
