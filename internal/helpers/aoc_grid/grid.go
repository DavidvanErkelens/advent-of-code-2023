package aoc_grid

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"fmt"
)

type Grid[T comparable] struct {
	Width      int
	Height     int
	dataPoints [][]T
}

type GridPoint[T comparable] struct {
	X     int
	Y     int
	Value T
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) GetTranslation() (int, int) {
	switch d {
	case Up:
		return 0, -1
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	case Right:
		return 1, 0
	default:
		panic("unhandled default case")
	}
}

func NewRuneGrid(input string) Grid[rune] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	dataPoints := make([][]rune, 0)

	for _, line := range lines {
		dataPoints = append(dataPoints, []rune(line))
	}

	return Grid[rune]{
		Width:      width,
		Height:     height,
		dataPoints: dataPoints,
	}
}

func NewEmptyRuneGrid(width, height int, initialValue rune) Grid[rune] {
	dataPoints := make([][]rune, height)
	for i := 0; i < width; i++ {
		dataPoints[i] = make([]rune, width)
		for j := range dataPoints[i] {
			dataPoints[i][j] = initialValue
		}
	}
	return Grid[rune]{
		Width:      width,
		Height:     height,
		dataPoints: dataPoints,
	}
}

func NewRuneGridFromParsedData(input [][]rune) Grid[rune] {
	width := len(input[0])
	height := len(input)

	return Grid[rune]{
		Width:      width,
		Height:     height,
		dataPoints: input,
	}
}

func NewIntGrid(input string, separator string) Grid[int] {
	lines := helpers.SplitLines(input)
	width := len(lines[0])
	height := len(lines)

	dataPoints := make([][]int, 0)

	for _, line := range lines {
		numbers := helpers.StringListOfNumericValuesToSlice(line, separator)
		dataPoints = append(dataPoints, numbers)
	}

	return Grid[int]{
		Width:      width,
		Height:     height,
		dataPoints: dataPoints,
	}
}

func NewEmptyIntGrid(width, height, initialValue int) Grid[int] {
	dataPoints := make([][]int, height)
	for i := 0; i < width; i++ {
		dataPoints[i] = make([]int, width)
		for j := range dataPoints[i] {
			dataPoints[i][j] = initialValue
		}
	}
	return Grid[int]{
		Width:      width,
		Height:     height,
		dataPoints: dataPoints,
	}
}

func (g *Grid[T]) Lines() [][]T {
	return g.dataPoints
}

func (g *Grid[T]) SetValueAt(x, y int, val T) {
	g.dataPoints[y][x] = val
}

func (g *Grid[T]) ValueAt(x int, y int) T {
	return g.dataPoints[y][x]
}

func (g *Grid[T]) ValueAtPoint(point aoc_point.Point) T {
	return g.dataPoints[point.Y][point.X]
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

func (g *Grid[T]) PointInRange(point aoc_point.Point) bool {
	return g.InRange(point.X, point.Y)
}

func (g *Grid[T]) FindLocationOfValue(val T) (int, int) {
	for y, line := range g.dataPoints {
		for x, value := range line {
			if value == val {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g *Grid[T]) FindAllLocationsOfValue(val T) []GridPoint[T] {
	points := make([]GridPoint[T], 0)
	for y, line := range g.dataPoints {
		for x, value := range line {
			if value == val {
				points = append(points, GridPoint[T]{
					X:     x,
					Y:     y,
					Value: value,
				})
			}
		}
	}

	return points
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
			neighbors = append(neighbors, GridPoint[T]{
				X:     p.x,
				Y:     p.y,
				Value: g.dataPoints[p.y][p.x],
			})
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

func (g *Grid[T]) PrintIndexFn(stringify func(T, int, int) string) {
	for y, line := range g.Lines() {
		for x, value := range line {
			fmt.Print(stringify(value, x, y))
		}
		fmt.Print("\n")
	}
}

func (g *Grid[T]) AsSingleLineString(stringify func(T) string) string {
	result := ""
	for _, line := range g.Lines() {
		for _, value := range line {
			result += stringify(value)
		}
	}
	return result
}

func (g *Grid[T]) InsertRowWithValue(at int, value T) {
	valuesToInsert := make([]T, 0)

	for i := 0; i < len(g.dataPoints[0]); i++ {
		valuesToInsert = append(valuesToInsert, value)
	}

	g.dataPoints = helpers.Insert(g.dataPoints, at, valuesToInsert)
	g.Height += 1
}

func (g *Grid[T]) InsertColumnWithValue(at int, value T) {
	for i := 0; i < len(g.dataPoints); i++ {
		g.dataPoints[i] = helpers.Insert(g.dataPoints[i], at, value)
	}
	g.Width += 1
}

func (g *Grid[T]) RowsHaveSameContents(rowOne, rowTwo int) bool {
	for i := 0; i < g.Width; i++ {
		if g.dataPoints[rowOne][i] != g.dataPoints[rowTwo][i] {
			return false
		}
	}

	return true
}

func (g *Grid[T]) ColumnsHaveSameContents(colOne, colTwo int) bool {
	for i := 0; i < g.Height; i++ {
		if g.dataPoints[i][colOne] != g.dataPoints[i][colTwo] {
			return false
		}
	}

	return true
}
func (g *Grid[T]) DifferenceBetweenRows(rowOne, rowTwo int) int {
	difference := 0
	for i := 0; i < g.Width; i++ {
		if g.dataPoints[rowOne][i] != g.dataPoints[rowTwo][i] {
			difference += 1
		}
	}

	return difference
}

func (g *Grid[T]) DifferenceBetweenColumns(colOne, colTwo int) int {
	difference := 0
	for i := 0; i < g.Height; i++ {
		if g.dataPoints[i][colOne] != g.dataPoints[i][colTwo] {
			difference += 1
		}
	}

	return difference
}
