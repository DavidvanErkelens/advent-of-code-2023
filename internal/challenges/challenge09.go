package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"strconv"
)

func NewChallenge09() Challenge09 {
	return Challenge09{}
}

type Challenge09 struct {
}

func isAllZeroes(list []int) bool {
	for _, i := range list {
		if i != 0 {
			return false
		}
	}
	return true
}

func (c Challenge09) parseLine(line string) [][]int {
	numbers := helpers.StringListOfNumericValuesToSlice(line, " ")
	sequences := [][]int{numbers}

	currentSequence := numbers
	for !isAllZeroes(currentSequence) {
		newSequence := make([]int, 0, len(currentSequence)-1)
		for i := 1; i < len(currentSequence); i++ {
			newSequence = append(newSequence, currentSequence[i]-currentSequence[i-1])
		}
		sequences = append(sequences, newSequence)
		currentSequence = newSequence
	}

	return sequences
}

func (c Challenge09) RunPartOne(input string) string {
	lines := helpers.SplitLines(input)
	lastValues := make([]int, 0)

	for _, line := range lines {
		sequences := c.parseLine(line)

		for i := len(sequences) - 2; i >= 0; i-- {
			lastElement := helpers.GetLastElement(sequences[i])
			newElement := lastElement + helpers.GetLastElement(sequences[i+1])
			sequences[i] = append(sequences[i], newElement)
		}

		lastValues = append(lastValues, helpers.GetLastElement(sequences[0]))
	}

	return strconv.Itoa(aoc_math.Sum(lastValues))
}

func (c Challenge09) RunPartTwo(input string) string {
	lines := helpers.SplitLines(input)
	firstValues := make([]int, 0)

	for _, line := range lines {
		sequences := c.parseLine(line)

		for i := len(sequences) - 2; i >= 0; i-- {
			firstElement := sequences[i][0]
			newElement := firstElement - sequences[i+1][0]
			sequences[i] = append([]int{newElement}, sequences[i]...)
		}

		firstValues = append(firstValues, sequences[0][0])
	}

	return strconv.Itoa(aoc_math.Sum(firstValues))
}

func (c Challenge09) DataFolder() string {
	return "09"
}
