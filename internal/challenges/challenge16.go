package challenges

import (
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"strconv"
)

func NewChallenge16() Challenge16 {
	return Challenge16{}
}

type Challenge16 struct {
}

type beam struct {
	location  aoc_point.Point
	direction aoc_grid.Direction
}

func (b *beam) nextStep() {
	b.location.ApplyTranslation(b.direction.GetTranslation())
}

func getEnergizedTilesForStartPosition(grid aoc_grid.Grid[rune], start aoc_point.Point, direction aoc_grid.Direction) int {
	energizedGrid := aoc_grid.NewEmptyIntGrid(grid.Width, grid.Height)
	energizedValues := 1

	seen := make(map[string]interface{})

	beams := []beam{{location: start, direction: direction}}
	energizedGrid.SetValueAt(start.X, start.Y, 1)

	for len(beams) > 0 {
		b := beams[0]
		beams = beams[1:]

		key := string(rune(b.location.X)) + "-" + string(rune(b.location.Y)) + "-" + string(rune(b.direction))
		if _, ok := seen[key]; ok {
			continue
		}

		newBeams := getNewBeams(grid, b)
		for _, nb := range newBeams {
			if energizedGrid.ValueAt(nb.location.X, nb.location.Y) == 0 {
				energizedValues += 1
			}
			energizedGrid.SetValueAt(nb.location.X, nb.location.Y, 1)
		}

		beams = append(beams, newBeams...)
		seen[key] = struct{}{}
	}

	//energizedGrid.Print(func(i int) string {
	//	if i == 0 {
	//		return "."
	//	}
	//	return "#"
	//})

	return energizedValues
}

func getNewBeams(grid aoc_grid.Grid[rune], b beam) []beam {
	result := make([]beam, 0)

	newDirections := getNewDirections(b.direction, grid.ValueAt(b.location.X, b.location.Y))

	for _, dir := range newDirections {
		nb := beam{
			location:  b.location,
			direction: dir,
		}

		nb.nextStep()

		if grid.InRange(nb.location.X, nb.location.Y) {
			result = append(result, nb)
		}
	}
	return result
}

func getNewDirections(direction aoc_grid.Direction, value rune) []aoc_grid.Direction {
	switch {
	case value == '.',
		value == '-' && (direction == aoc_grid.Left || direction == aoc_grid.Right),
		value == '|' && (direction == aoc_grid.Up || direction == aoc_grid.Down):
		return []aoc_grid.Direction{direction}

	case value == '/' && direction == aoc_grid.Up,
		value == '\\' && direction == aoc_grid.Down:
		return []aoc_grid.Direction{aoc_grid.Right}

	case value == '/' && direction == aoc_grid.Down, value == '\\' && direction == aoc_grid.Up:
		return []aoc_grid.Direction{aoc_grid.Left}

	case value == '/' && direction == aoc_grid.Right, value == '\\' && direction == aoc_grid.Left:
		return []aoc_grid.Direction{aoc_grid.Up}

	case value == '/' && direction == aoc_grid.Left, value == '\\' && direction == aoc_grid.Right:
		return []aoc_grid.Direction{aoc_grid.Down}

	case value == '-' && (direction == aoc_grid.Up || direction == aoc_grid.Down):
		return []aoc_grid.Direction{aoc_grid.Left, aoc_grid.Right}

	case value == '|' && (direction == aoc_grid.Left || direction == aoc_grid.Right):
		return []aoc_grid.Direction{aoc_grid.Up, aoc_grid.Down}
	}
	panic("not reachable")
}

func (c Challenge16) RunPartOne(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	energizedValues := getEnergizedTilesForStartPosition(grid, aoc_point.NewPoint(0, 0), aoc_grid.Right)

	return strconv.Itoa(energizedValues)
}

func (c Challenge16) RunPartTwo(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	maxValue := 0

	for i := 0; i < grid.Width; i++ {
		downValue := getEnergizedTilesForStartPosition(grid, aoc_point.NewPoint(i, 0), aoc_grid.Down)
		//fmt.Println()
		upValue := getEnergizedTilesForStartPosition(grid, aoc_point.NewPoint(i, grid.Height-1), aoc_grid.Up)
		//fmt.Println()
		maxValue = max(maxValue, upValue, downValue)
	}

	for i := 0; i < grid.Height; i++ {
		leftValue := getEnergizedTilesForStartPosition(grid, aoc_point.NewPoint(0, i), aoc_grid.Right)
		//fmt.Println()
		rightValue := getEnergizedTilesForStartPosition(grid, aoc_point.NewPoint(grid.Width-1, i), aoc_grid.Left)
		//fmt.Println()
		maxValue = max(maxValue, rightValue, leftValue)
	}

	return strconv.Itoa(maxValue)
}

func (c Challenge16) DataFolder() string {
	return "16"
}
