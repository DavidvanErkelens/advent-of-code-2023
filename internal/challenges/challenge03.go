package challenges

import (
	"advent-of-code-2023/internal/helpers/aoc_grid"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"advent-of-code-2023/internal/helpers/aoc_range"
	"fmt"
	"strconv"
	"unicode"
)

func NewChallenge03() Challenge03 {
	return Challenge03{}
}

type Challenge03 struct {
}

func (c Challenge03) RunPartOne(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	partNumbers := make([]int, 0)

	for lineIdx, line := range grid.Lines() {
		startIdx := -1
		prevWasDigit := false

		handleEndOfDigit := func(currentIdx int) {
			digit := line[startIdx:currentIdx]

			xRange := aoc_range.NewRangeSlice(startIdx-1, currentIdx)
			yRange := aoc_range.NewRangeSlice(lineIdx-1, lineIdx+1)

			adjacentToSymbol := false
			for _, x := range xRange {
				for _, y := range yRange {
					if !grid.InRange(x, y) {
						continue
					}

					value := grid.ValueAt(x, y)
					if !unicode.IsDigit(value) && value != '.' {
						adjacentToSymbol = true
						break
					}
				}

				if adjacentToSymbol {
					break
				}
			}

			if adjacentToSymbol {
				digitValue, _ := strconv.Atoi(string(digit))
				partNumbers = append(partNumbers, digitValue)
			}
		}

		for currentIdx, char := range line {
			if unicode.IsDigit(char) {
				if !prevWasDigit {
					startIdx = currentIdx
				}

				prevWasDigit = true
			} else {
				if prevWasDigit {
					handleEndOfDigit(currentIdx)
				}

				prevWasDigit = false
			}
		}

		if prevWasDigit {
			handleEndOfDigit(grid.Width)
		}
	}

	return strconv.Itoa(aoc_math.Sum(partNumbers))
}

func (c Challenge03) RunPartTwo(input string) string {
	grid := aoc_grid.NewRuneGrid(input)
	adjacentToGears := make(map[string][]int, 0)

	for lineIdx, line := range grid.Lines() {
		startIdx := -1
		prevWasDigit := false

		handleEndOfDigit := func(currentIdx int) {
			digit := line[startIdx:currentIdx]

			xRange := aoc_range.NewRangeSlice(startIdx-1, currentIdx)
			yRange := aoc_range.NewRangeSlice(lineIdx-1, lineIdx+1)

			adjacentToSymbol := false
			for _, x := range xRange {
				for _, y := range yRange {
					if !grid.InRange(x, y) {
						continue
					}

					value := grid.ValueAt(x, y)
					if value == '*' {
						key := fmt.Sprintf("%d-%d", x, y)
						val, ok := adjacentToGears[key]
						digitValue, _ := strconv.Atoi(string(digit))

						if ok {
							adjacentToGears[key] = append(val, digitValue)
						} else {
							adjacentToGears[key] = []int{digitValue}
						}

						adjacentToSymbol = true
						break
					}
				}

				if adjacentToSymbol {
					break
				}
			}
		}

		for currentIdx, char := range line {
			if unicode.IsDigit(char) {
				if !prevWasDigit {
					startIdx = currentIdx
				}

				prevWasDigit = true
			} else {
				if prevWasDigit {
					handleEndOfDigit(currentIdx)
				}

				prevWasDigit = false
			}
		}

		if prevWasDigit {
			handleEndOfDigit(grid.Width)
		}
	}

	products := make([]int, 0)
	for _, adjacentParts := range adjacentToGears {
		if len(adjacentParts) == 2 {
			products = append(products, adjacentParts[0]*adjacentParts[1])
		}
	}

	return strconv.Itoa(aoc_math.Sum(products))
}

func (c Challenge03) DataFolder() string {
	return "03"
}
