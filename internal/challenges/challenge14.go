package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_loop"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"slices"
	"strconv"
)

func NewChallenge14() Challenge14 {
	return Challenge14{}
}

type Challenge14 struct {
}

const (
	roundRock rune = 'O'
	cubeRock  rune = '#'
	empty     rune = '.'
)

const (
	up int = iota
	down
	left
	right
)

func (c Challenge14) RunPartOne(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	allRocks := helpers.Map(grid.FindAllLocationsOfValue(roundRock), func(p aoc_grid.GridPoint[rune]) aoc_point.Point {
		return aoc_point.NewPoint(p.X, p.Y)
	})

	for i, r := range allRocks {
		allRocks[i] = c.moveRock(&grid, r, 0, -1)
	}

	return strconv.Itoa(c.calculateResult(grid.Height, allRocks))
}

func (c Challenge14) RunPartTwo(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	requiredSteps := 1000000000
	allRocks := helpers.Map(grid.FindAllLocationsOfValue(roundRock), func(p aoc_grid.GridPoint[rune]) aoc_point.Point {
		return aoc_point.NewPoint(p.X, p.Y)
	})

	states := make(map[string]int)
	translations := map[int]struct{ dx, dy int }{
		up:    {dx: 0, dy: -1},
		down:  {dx: 0, dy: 1},
		left:  {dx: -1, dy: 0},
		right: {dx: 1, dy: 0},
	}

	for i := 1; i <= requiredSteps; i++ {
		for _, direction := range []int{up, left, down, right} {
			translation := translations[direction]
			allRocks = c.sortForDirection(allRocks, direction)
			for i, r := range allRocks {
				allRocks[i] = c.moveRock(&grid, r, translation.dx, translation.dy)
			}
		}

		state := grid.AsSingleLineString(func(r rune) string { return string(r) })

		if val, ok := states[state]; ok {
			loop := aoc_loop.NewLoop(val, i)
			i = loop.FindLastEqual(requiredSteps)
		}

		states[state] = i
	}

	return strconv.Itoa(c.calculateResult(grid.Height, allRocks))
}

func (c Challenge14) sortForDirection(rocks []aoc_point.Point, direction int) []aoc_point.Point {
	sortFunc := func(a, b aoc_point.Point) int {
		if direction == up {
			return a.Y - b.Y
		}
		if direction == down {
			return b.Y - a.Y
		}
		if direction == left {
			return a.X - b.X
		}
		if direction == right {
			return b.X - a.X
		}
		panic("not reachable")
	}

	slices.SortFunc(rocks, sortFunc)
	return rocks
}

func (c Challenge14) moveRock(grid *aoc_grid.Grid[rune], point aoc_point.Point, dx, dy int) aoc_point.Point {
	for grid.InRange(point.X+dx, point.Y+dy) && grid.ValueAt(point.X+dx, point.Y+dy) == empty {
		grid.SetValueAt(point.X+dx, point.Y+dy, roundRock)
		grid.SetValueAt(point.X, point.Y, empty)
		point.X += dx
		point.Y += dy
	}
	return point
}

func (c Challenge14) calculateResult(gridHeight int, rocks []aoc_point.Point) int {
	return helpers.Reduce(rocks, func(r aoc_point.Point, total int) int {
		return total + gridHeight - r.Y
	}, 0)
}

func (c Challenge14) DataFolder() string {
	return "14"
}
