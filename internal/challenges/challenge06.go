package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_assert"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func NewChallenge06() Challenge06 {
	return Challenge06{}
}

type Challenge06 struct {
}

func (c Challenge06) getUpperLowerBound(time, distance int) (int, int) {
	D := time*time - 4*distance
	lowerBound := (-1*float64(time) + math.Sqrt(float64(D))) / -2
	upperBound := (-1*float64(time) - math.Sqrt(float64(D))) / -2

	if lowerBound == math.Trunc(lowerBound) {
		lowerBound += 1
	}

	if upperBound == math.Trunc(upperBound) {
		lowerBound += 1
	}

	return int(math.Ceil(lowerBound)), int(math.Floor(upperBound))
}

func (c Challenge06) RunPartOne(input string) string {
	lines := helpers.SplitLines(input)
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")

	times := helpers.StringListOfNumericValuesToSlice(timeLine[1], " ")
	distances := helpers.StringListOfNumericValuesToSlice(distanceLine[1], " ")

	aoc_assert.Assert(len(times) == len(distances), "different lengths")

	outcome := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		lowerbound, upperbound := c.getUpperLowerBound(time, distance)
		outcome *= upperbound - lowerbound + 1
	}

	return strconv.Itoa(outcome)
}

func (c Challenge06) RunPartOneNaive(input string) string {
	lines := helpers.SplitLines(input)
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")

	times := helpers.StringListOfNumericValuesToSlice(timeLine[1], " ")
	distances := helpers.StringListOfNumericValuesToSlice(distanceLine[1], " ")

	winningSetups := make([]int, 0)

	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		waysToBeat := 0

		for j := 1; j < time; j++ {
			travelTime := time - j
			travelDistance := travelTime * j

			if travelDistance > distance {
				waysToBeat += 1
			}
		}

		winningSetups = append(winningSetups, waysToBeat)
	}

	return strconv.Itoa(aoc_math.Product(winningSetups))
}

func (c Challenge06) RunPartTwo(input string) string {
	lines := helpers.SplitLines(input)
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")

	timeRunes := []rune(timeLine[1])
	timeRunes = helpers.Filter(timeRunes, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
	time := helpers.NaiveStringToInt(string(timeRunes))

	distanceRunes := []rune(distanceLine[1])
	distanceRunes = helpers.Filter(distanceRunes, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
	distance := helpers.NaiveStringToInt(string(distanceRunes))

	lowerbound, upperbound := c.getUpperLowerBound(time, distance)
	outcome := upperbound - lowerbound + 1

	return strconv.Itoa(outcome)
}

func (c Challenge06) RunPartTwoNaive(input string) string {
	lines := helpers.SplitLines(input)
	timeLine := strings.Split(lines[0], ":")
	distanceLine := strings.Split(lines[1], ":")

	timeRunes := []rune(timeLine[1])
	timeRunes = helpers.Filter(timeRunes, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
	time := helpers.NaiveStringToInt(string(timeRunes))

	distanceRunes := []rune(distanceLine[1])
	distanceRunes = helpers.Filter(distanceRunes, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
	distance := helpers.NaiveStringToInt(string(distanceRunes))

	winningStart := 0
	winningEnd := time

	for i := 1; i < time; i++ {
		travelTime := time - i
		travelDistance := travelTime * i

		if travelDistance > distance {
			winningStart = i
			break
		}
	}

	for i := time; i > 0; i-- {
		travelTime := time - i
		travelDistance := travelTime * i

		if travelDistance > distance {
			winningEnd = i
			break
		}
	}

	return strconv.Itoa(winningEnd - winningStart + 1)
}

func (c Challenge06) DataFolder() string {
	return "06"
}
