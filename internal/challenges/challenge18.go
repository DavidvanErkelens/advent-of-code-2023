package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"strconv"
	"strings"
)

func NewChallenge18() Challenge18 {
	return Challenge18{}
}

type Challenge18 struct {
}

type inputElement struct {
	direction aoc_grid.Direction
	value     int
}

func parseInput(input string) []inputElement {
	lines := helpers.SplitLines(input)
	inputs := make([]inputElement, 0, len(lines))
	for _, line := range lines {
		elements := strings.Split(line, " ")
		inputs = append(inputs, inputElement{
			direction: inputToDirection(elements[0]),
			value:     helpers.NaiveStringToInt(elements[1]),
		})
	}

	return inputs
}

func inputToDirection(input string) aoc_grid.Direction {
	switch input {
	case "R":
		return aoc_grid.Right
	case "L":
		return aoc_grid.Left
	case "U":
		return aoc_grid.Up
	case "D":
		return aoc_grid.Down
	}
	panic("not reachable")
}

func intToDir(val int) aoc_grid.Direction {
	switch val {
	case 0:
		return aoc_grid.Right
	case 1:
		return aoc_grid.Down
	case 2:
		return aoc_grid.Left
	case 3:
		return aoc_grid.Up
	}
	panic("not reachable")
}

func parseInputTwo(input string) []inputElement {
	lines := helpers.SplitLines(input)
	inputs := make([]inputElement, 0, len(lines))
	for _, line := range lines {
		elements := strings.Split(line, " ")
		dir, val := parseHexCode(elements[2][2:9])
		inputs = append(inputs, inputElement{
			direction: dir,
			value:     val,
		})
	}

	return inputs
}

func parseHexCode(input string) (aoc_grid.Direction, int) {
	value, _ := strconv.ParseInt(input[:5], 16, 32)
	direction := intToDir(helpers.NaiveStringToInt(string(input[5])))
	return direction, int(value)
}

func (c Challenge18) RunPartOne(input string) string {
	instructions := parseInput(input)

	pos := aoc_point.NewPoint(0, 0)

	points := make([]aoc_point.Point, len(instructions)+1)
	points[0] = pos

	steps := 0

	for i, instr := range instructions {
		dx, dy := instr.direction.GetTranslation()
		pos.ApplyTranslation(dx*(instr.value), dy*(instr.value))
		points[i+1] = pos
		steps += instr.value
	}

	sum := 0
	for i := 0; i < len(points)-1; i++ {
		sum += (points[i+1].X + points[i].X) * (points[i+1].Y - points[i].Y)
	}

	abs := aoc_math.Abs(sum)

	insidePoints := abs / 2
	outSidePoints := steps

	totalPoints := insidePoints + (outSidePoints / 2) + 1

	return strconv.Itoa(totalPoints)
}

func (c Challenge18) RunPartTwo(input string) string {
	instructions := parseInputTwo(input)

	pos := aoc_point.NewPoint(0, 0)

	points := make([]aoc_point.Point, len(instructions)+1)
	points[0] = pos

	steps := 0

	for i, instr := range instructions {
		dx, dy := instr.direction.GetTranslation()
		pos.ApplyTranslation(dx*(instr.value), dy*(instr.value))
		points[i+1] = pos
		steps += instr.value
	}

	sum := 0
	for i := 0; i < len(points)-1; i++ {
		sum += (points[i+1].X + points[i].X) * (points[i+1].Y - points[i].Y)
	}

	abs := aoc_math.Abs(sum)

	insidePoints := abs / 2
	outSidePoints := steps

	totalPoints := insidePoints + (outSidePoints / 2) + 1

	return strconv.Itoa(totalPoints)
}

func (c Challenge18) DataFolder() string {
	return "18"
}
