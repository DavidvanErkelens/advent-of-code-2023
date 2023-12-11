package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"strconv"
)

const (
	VerticalPipe       = '|'
	HorizontalPipe     = '-'
	NorthEastConnector = 'L'
	NorthWestConnector = 'J'
	SouthWestConnector = '7'
	SouthEastConnector = 'F'
	Ground             = '.'
	Start              = 'S'
	Water              = '#'
)

const (
	Up int = iota
	Down
	Left
	Right
)

func NewChallenge10() Challenge10 {
	return Challenge10{
		availableDirections: map[rune][]int{
			Start:              {Up, Down, Left, Right},
			VerticalPipe:       {Up, Down},
			HorizontalPipe:     {Left, Right},
			NorthWestConnector: {Up, Left},
			NorthEastConnector: {Up, Right},
			SouthWestConnector: {Down, Left},
			SouthEastConnector: {Down, Right},
		},
		availableConnectors: map[int][]rune{
			Up:    {VerticalPipe, SouthWestConnector, SouthEastConnector},
			Down:  {VerticalPipe, NorthWestConnector, NorthEastConnector},
			Left:  {HorizontalPipe, NorthEastConnector, SouthEastConnector},
			Right: {HorizontalPipe, NorthWestConnector, SouthWestConnector},
		},
		translations: map[int]struct{ x, y int }{
			Up:    {x: 0, y: -1},
			Down:  {x: 0, y: 1},
			Left:  {x: -1, y: 0},
			Right: {x: 1, y: 0},
		},
		betterSymbols: map[rune]rune{
			Start:              'S',
			VerticalPipe:       '┃',
			HorizontalPipe:     '━',
			NorthWestConnector: '┛',
			NorthEastConnector: '┗',
			SouthWestConnector: '┓',
			SouthEastConnector: '┏',
			Ground:             '.',
			Water:              '#',
		},
	}
}

type Challenge10 struct {
	availableDirections map[rune][]int
	availableConnectors map[int][]rune
	translations        map[int]struct{ x, y int }
	betterSymbols       map[rune]rune
}

func (c Challenge10) RunPartOne(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	sx, sy := grid.FindLocationOfValue(Start)
	toCheck := []struct{ x, y, length int }{
		{x: sx, y: sy, length: 0},
	}
	alreadyFound := make(map[string]interface{}, 0)

	maxDistance := 0

	for len(toCheck) > 0 {
		pos := toCheck[0]
		pipeAtPos := grid.ValueAt(pos.x, pos.y)

		directions := c.availableDirections[pipeAtPos]

		for _, dir := range directions {
			nextDestX := pos.x + c.translations[dir].x
			nextDestY := pos.y + c.translations[dir].y

			if !grid.InRange(nextDestX, nextDestY) {
				continue
			}

			nextDest := grid.ValueAt(nextDestX, nextDestY)

			if !helpers.ContainsElement(c.availableConnectors[dir], nextDest) {
				continue
			}

			key := string(rune(nextDestX)) + "-" + string(rune(nextDestY))
			if _, contains := alreadyFound[key]; contains {
				continue
			}

			distance := pos.length + 1
			if distance > maxDistance {
				maxDistance = distance
			}

			alreadyFound[key] = distance
			toCheck = append(toCheck, struct{ x, y, length int }{x: nextDestX, y: nextDestY, length: distance})
		}

		toCheck = toCheck[1:]
	}

	return strconv.Itoa(maxDistance)
}

func (c Challenge10) expand(grid aoc_grid.Grid[rune]) aoc_grid.Grid[rune] {
	emptyLine := make([]rune, 0)

	for i := 0; i < len(grid.Lines()[0])*2+2; i++ {
		emptyLine = append(emptyLine, Ground)
	}

	newGridData := [][]rune{emptyLine}

	for _, line := range grid.Lines() {
		newLineOne := []rune{Ground}
		newLineTwo := []rune{Ground}

		for _, value := range line {
			switch value {
			case HorizontalPipe:
				newLineOne = append(newLineOne, HorizontalPipe, HorizontalPipe)
				newLineTwo = append(newLineTwo, Ground, Ground)
			case VerticalPipe:
				newLineOne = append(newLineOne, VerticalPipe, Ground)
				newLineTwo = append(newLineTwo, VerticalPipe, Ground)
			case NorthWestConnector:
				newLineOne = append(newLineOne, NorthWestConnector, Ground)
				newLineTwo = append(newLineTwo, Ground, Ground)
			case NorthEastConnector:
				newLineOne = append(newLineOne, NorthEastConnector, HorizontalPipe)
				newLineTwo = append(newLineTwo, Ground, Ground)
			case SouthWestConnector:
				newLineOne = append(newLineOne, SouthWestConnector, Ground)
				newLineTwo = append(newLineTwo, VerticalPipe, Ground)
			case SouthEastConnector:
				newLineOne = append(newLineOne, SouthEastConnector, HorizontalPipe)
				newLineTwo = append(newLineTwo, VerticalPipe, Ground)
			case Ground:
				newLineOne = append(newLineOne, Ground, Ground)
				newLineTwo = append(newLineTwo, Ground, Ground)
			case Start:
				newLineOne = append(newLineOne, Start, Start)
				newLineTwo = append(newLineTwo, Start, Start)
			}
		}

		newLineOne = append(newLineOne, Ground)
		newLineTwo = append(newLineTwo, Ground)

		newGridData = append(newGridData, newLineOne, newLineTwo)
	}

	newGridData = append(newGridData, emptyLine)

	return aoc_grid.NewRuneGridFromParsedData(newGridData)
}

func (c Challenge10) floodGrid(grid aoc_grid.Grid[rune]) {
	toRun := []struct{ x, y int }{{0, 0}}

	for len(toRun) > 0 {
		val := toRun[0]
		toRun = toRun[1:]

		if grid.ValueAt(val.x, val.y) != Ground {
			continue
		}

		grid.SetValueAt(val.x, val.y, Water)

		neighbors := grid.GetNeighbors(val.x, val.y)
		for _, n := range neighbors {
			toRun = append(toRun, struct{ x, y int }{x: n.X, y: n.Y})
		}
	}
}

func (c Challenge10) RunPartTwo(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	sx, sy := grid.FindLocationOfValue(Start)
	toCheck := []struct{ x, y int }{
		{x: sx, y: sy},
	}
	alreadyFound := make(map[string]interface{})

	for len(toCheck) > 0 {
		pos := toCheck[0]
		pipeAtPos := grid.ValueAt(pos.x, pos.y)

		directions := c.availableDirections[pipeAtPos]

		for _, dir := range directions {
			nextDestX := pos.x + c.translations[dir].x
			nextDestY := pos.y + c.translations[dir].y

			if !grid.InRange(nextDestX, nextDestY) {
				continue
			}

			nextDest := grid.ValueAt(nextDestX, nextDestY)

			if !helpers.ContainsElement(c.availableConnectors[dir], nextDest) {
				continue
			}

			key := string(rune(nextDestX)) + "-" + string(rune(nextDestY))
			if _, contains := alreadyFound[key]; contains {
				continue
			}

			alreadyFound[key] = 0
			toCheck = append(toCheck, struct{ x, y int }{x: nextDestX, y: nextDestY})
		}

		toCheck = toCheck[1:]
	}

	for y, line := range grid.Lines() {
		for x, value := range line {
			key := string(rune(x)) + "-" + string(rune(y))
			if _, contains := alreadyFound[key]; contains {
			} else if value == 'S' {
			} else {
				grid.SetValueAt(x, y, Ground)
			}
		}
	}

	largerGrid := c.expand(grid)
	c.floodGrid(largerGrid)

	pointsInsideLoop := 0
	for y, line := range grid.Lines() {
		for x, value := range line {
			if value == Ground {
				if largerGrid.ValueAt(x*2+1, y*2+1) == Ground {
					pointsInsideLoop += 1
				}
			}
		}
	}

	return strconv.Itoa(pointsInsideLoop)
}

func (c Challenge10) DataFolder() string {
	return "10"
}
