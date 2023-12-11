package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"strconv"
)

const (
	Galaxy = '#'
	Empty  = '.'
	Filled = 'X'
)

func NewChallenge11() Challenge11 {
	return Challenge11{}
}

type Challenge11 struct {
}

func (c Challenge11) cosmicExpansion(g *aoc_grid.Grid[rune]) {
	numberOfRows := len(g.Lines())
	// Insert extra rows
	for i := 0; i < numberOfRows; i++ {
		if helpers.AllSatisfies(g.Lines()[i], func(r rune) bool { return r == Empty }) {
			g.InsertRowWithValue(i+1, Empty)
			i += 1
			numberOfRows += 1
		}
	}

	// Insert extra columns
	numberOfColumns := len(g.Lines()[0])
	for i := 0; i < numberOfColumns; i++ {
		allEmpty := true
		for j := 0; j < len(g.Lines()); j++ {
			if g.ValueAt(i, j) == Galaxy {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			g.InsertColumnWithValue(i, Empty)
			numberOfColumns += 1
			i += 1
		}
	}
}

func (c Challenge11) getEmptyRows(g aoc_grid.Grid[rune]) []int {
	emptyRows := make([]int, 0)
	for idx, line := range g.Lines() {
		if helpers.AllSatisfies(line, func(r rune) bool { return r == Empty }) {
			emptyRows = append(emptyRows, idx)
		}
	}
	return emptyRows
}

func (c Challenge11) getEmptyColumns(g aoc_grid.Grid[rune]) []int {
	emptyColumns := make([]int, 0)
	for i := 0; i < len(g.Lines()[0]); i++ {
		allEmpty := true
		for j := 0; j < len(g.Lines()); j++ {
			if g.ValueAt(i, j) == Galaxy {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			emptyColumns = append(emptyColumns, i)
		}
	}
	return emptyColumns
}

func (c Challenge11) RunPartOne(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	c.cosmicExpansion(&grid)

	galaxies := grid.FindAllLocationsOfValue(Galaxy)
	totalLengths := 0

	for i := 0; i < len(galaxies); i++ {
		galaxyOne := aoc_point.NewPoint(galaxies[i].X, galaxies[i].Y)

		for j := i + 1; j < len(galaxies); j++ {
			galaxyTwo := aoc_point.NewPoint(galaxies[j].X, galaxies[j].Y)
			totalLengths += galaxyOne.ManhattanDistance(galaxyTwo)
		}
	}

	return strconv.Itoa(totalLengths)
}

func (c Challenge11) RunPartTwo(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	galaxies := grid.FindAllLocationsOfValue(Galaxy)
	totalLengths := 0
	replacementFactor := 1000000 - 1
	// replacementFactor := 99

	emptyRows := c.getEmptyRows(grid)
	emptyColumns := c.getEmptyColumns(grid)

	for i := 0; i < len(galaxies); i++ {
		galaxyOne := aoc_point.NewPoint(galaxies[i].X, galaxies[i].Y)

		for j := i + 1; j < len(galaxies); j++ {
			galaxyTwo := aoc_point.NewPoint(galaxies[j].X, galaxies[j].Y)
			emptyRowsInBetween := len(helpers.Filter(emptyRows, func(i int) bool {
				return i > min(galaxyOne.Y, galaxies[j].Y) && i < max(galaxyOne.Y, galaxies[j].Y)
			}))
			emptyColsInBetween := len(helpers.Filter(emptyColumns, func(i int) bool {
				return i > min(galaxyOne.X, galaxies[j].X) && i < max(galaxyOne.X, galaxies[j].X)
			}))

			distanceWithoutExpansion := galaxyTwo.ManhattanDistance(galaxyOne)
			distance := distanceWithoutExpansion + emptyRowsInBetween*replacementFactor + emptyColsInBetween*replacementFactor
			totalLengths += distance
		}
	}

	return strconv.Itoa(totalLengths)
}

func (c Challenge11) DataFolder() string {
	return "11"
}
