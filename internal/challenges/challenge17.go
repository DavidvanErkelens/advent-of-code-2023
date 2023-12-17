package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_assert"
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_point"
	"math"
	"strconv"
)

func NewChallenge17() Challenge17 {
	return Challenge17{}
}

type Challenge17 struct {
}

type movingPath struct {
	location  aoc_point.Point
	cost      int
	straights int
	direction aoc_grid.Direction
}

func (c Challenge17) RunPartOne(input string) string {
	grid := aoc_grid.NewIntGrid(input, "")

	queue := helpers.BuildPriorityQueue(movingPath{
		location:  aoc_point.NewPoint(0, 0),
		cost:      0,
		straights: 0,
		direction: aoc_grid.Right,
	}, 0)

	smallestCosts := make(map[string]int)
	smallestEndCost := math.MaxInt

	allowedDirections := map[aoc_grid.Direction][]aoc_grid.Direction{
		aoc_grid.Up:    {aoc_grid.Left, aoc_grid.Up, aoc_grid.Right},
		aoc_grid.Down:  {aoc_grid.Left, aoc_grid.Down, aoc_grid.Right},
		aoc_grid.Left:  {aoc_grid.Up, aoc_grid.Left, aoc_grid.Down},
		aoc_grid.Right: {aoc_grid.Up, aoc_grid.Right, aoc_grid.Down},
	}

	for queue.HasElements() {
		p := queue.Pop()

		key := strconv.Itoa(p.location.X) + "-" + strconv.Itoa(p.location.Y) + "-" + strconv.Itoa(int(p.direction)) + "-" + strconv.Itoa(p.straights)

		if p.straights > 3 {
			continue
		}

		if value, ok := smallestCosts[key]; ok && value <= p.cost {
			continue
		}

		smallestCosts[key] = p.cost

		if p.location.X == grid.Width-1 && p.location.Y == grid.Height-1 {
			smallestEndCost = min(smallestEndCost, p.cost)
			continue
		}

		for _, dir := range allowedDirections[p.direction] {
			if dir == p.direction && p.straights > 3 {
				continue
			}

			newLocation := aoc_point.Point{
				X: p.location.X,
				Y: p.location.Y,
			}
			newLocation.ApplyTranslation(dir.GetTranslation())

			if !grid.PointInRange(newLocation) {
				continue
			}

			straights := 1

			if dir == p.direction {
				straights += p.straights
			}

			newKey := strconv.Itoa(newLocation.X) + "-" + strconv.Itoa(newLocation.Y) + "-" + strconv.Itoa(int(dir)) + "-" + strconv.Itoa(straights)

			if value, ok := smallestCosts[newKey]; ok && value <= p.cost {
				continue
			}

			newPath := movingPath{
				location:  newLocation,
				cost:      p.cost + grid.ValueAtPoint(newLocation),
				straights: straights,
				direction: dir,
			}

			aoc_assert.Assert(newLocation.ManhattanDistance(p.location) == 1, "larger distance")

			queue.Push(newPath, newPath.cost)
		}
	}

	return strconv.Itoa(smallestEndCost)
}

func (c Challenge17) RunPartTwo(input string) string {
	grid := aoc_grid.NewIntGrid(input, "")

	queue := helpers.BuildPriorityQueue(movingPath{
		location:  aoc_point.NewPoint(0, 0),
		cost:      0,
		straights: 0,
		direction: aoc_grid.Right,
	}, 0)

	smallestCosts := make(map[string]int)
	smallestEndCost := math.MaxInt

	allowedDirections := map[aoc_grid.Direction][]aoc_grid.Direction{
		aoc_grid.Up:    {aoc_grid.Left, aoc_grid.Up, aoc_grid.Right},
		aoc_grid.Down:  {aoc_grid.Left, aoc_grid.Down, aoc_grid.Right},
		aoc_grid.Left:  {aoc_grid.Up, aoc_grid.Left, aoc_grid.Down},
		aoc_grid.Right: {aoc_grid.Up, aoc_grid.Right, aoc_grid.Down},
	}

	for queue.HasElements() {
		p := queue.Pop()

		key := strconv.Itoa(p.location.X) + "-" + strconv.Itoa(p.location.Y) + "-" + strconv.Itoa(int(p.direction)) + "-" + strconv.Itoa(p.straights)

		if p.straights > 10 {
			continue
		}

		if value, ok := smallestCosts[key]; ok && value <= p.cost {
			continue
		}

		smallestCosts[key] = p.cost

		if p.location.X == grid.Width-1 && p.location.Y == grid.Height-1 {
			smallestEndCost = min(smallestEndCost, p.cost)
			continue
		}

		for _, dir := range allowedDirections[p.direction] {
			if dir != p.direction && p.straights < 4 {
				continue
			}

			if dir == p.direction && p.straights > 10 {
				continue
			}

			newLocation := aoc_point.Point{
				X: p.location.X,
				Y: p.location.Y,
			}
			newLocation.ApplyTranslation(dir.GetTranslation())

			if !grid.PointInRange(newLocation) {
				continue
			}

			straights := 1

			if dir == p.direction {
				straights += p.straights
			}

			newKey := strconv.Itoa(newLocation.X) + "-" + strconv.Itoa(newLocation.Y) + "-" + strconv.Itoa(int(dir)) + "-" + strconv.Itoa(straights)

			if value, ok := smallestCosts[newKey]; ok && value <= p.cost {
				continue
			}

			newPath := movingPath{
				location:  newLocation,
				cost:      p.cost + grid.ValueAtPoint(newLocation),
				straights: straights,
				direction: dir,
			}

			aoc_assert.Assert(newLocation.ManhattanDistance(p.location) == 1, "larger distance")

			queue.Push(newPath, newPath.cost)
		}
	}

	return strconv.Itoa(smallestEndCost)
}

func (c Challenge17) DataFolder() string {
	return "17"
}
