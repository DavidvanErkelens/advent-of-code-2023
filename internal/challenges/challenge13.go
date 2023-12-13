package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"strconv"
)

func NewChallenge13() Challenge13 {
	return Challenge13{}
}

type Challenge13 struct {
}

func (c Challenge13) RunPartOne(input string) string {
	return c.solve(input, 0)
}

func (c Challenge13) RunPartTwo(input string) string {
	return c.solve(input, 1)
}

func (c Challenge13) solve(input string, difference int) string {
	grids := helpers.SplitParagraphs(input)
	numberOfRows := 0
	numberOfColumns := 0

	for _, g := range grids {
		grid := aoc_grid.NewRuneGrid(g)
		verticalLine := c.getVerticalReflection(grid, difference)
		horizontalLine := c.getHorizontalReflection(grid, difference)

		if verticalLine > 0 {
			numberOfColumns += verticalLine
		}
		if horizontalLine > 0 {
			numberOfRows += horizontalLine
		}
	}
	return strconv.Itoa(100*numberOfRows + numberOfColumns)
}

func (c Challenge13) getVerticalReflection(grid aoc_grid.Grid[rune], difference int) int {
	for i := 0; i < grid.Width-1; i++ {
		runningDifference := grid.DifferenceBetweenColumns(i, i+1)

		if runningDifference > difference {
			continue
		}

		for j := 1; j <= i; j++ {
			if i+j+1 >= grid.Width {
				break
			}
			runningDifference += grid.DifferenceBetweenColumns(i-j, i+j+1)
			if runningDifference > difference {
				break
			}
		}

		if runningDifference == difference {
			return i + 1
		}
	}

	return -1
}

func (c Challenge13) getHorizontalReflection(grid aoc_grid.Grid[rune], difference int) int {
	for i := 0; i < grid.Height-1; i++ {
		runningDifference := grid.DifferenceBetweenRows(i, i+1)

		if runningDifference > difference {
			continue
		}

		for j := 1; j <= i; j++ {
			if i+j+1 >= grid.Height {
				break
			}
			runningDifference += grid.DifferenceBetweenRows(i-j, i+j+1)
			if runningDifference > difference {
				break
			}
		}

		if runningDifference == difference {
			return i + 1
		}
	}

	return -1
}

func (c Challenge13) DataFolder() string {
	return "13"
}
