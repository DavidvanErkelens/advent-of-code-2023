package challenges

import (
	"advent-of-code-2023/internal/common"
	"fmt"
	"strconv"
	"unicode"
)

func NewChallengeThree() ChallengeThree {
	return ChallengeThree{}
}

type ChallengeThree struct {
}

func (c ChallengeThree) RunPartOne(input string) string {
	grid := common.GridFromInput(input)
	partNumbers := make([]int, 0)

	for lineIdx, line := range grid.Lines() {
		startIdx := -1
		prevWasDigit := false

		handleEndOfDigit := func(currentIdx int) {
			digit := line[startIdx:currentIdx]

			xRange := common.NewSlice(startIdx-1, currentIdx, 1)
			yRange := common.NewSlice(lineIdx-1, lineIdx+1, 1)

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

	return strconv.Itoa(common.Sum(partNumbers))
}

func (c ChallengeThree) RunPartTwo(input string) string {
	grid := common.GridFromInput(input)
	adjacentToGears := make(map[string][]int, 0)

	for lineIdx, line := range grid.Lines() {
		startIdx := -1
		prevWasDigit := false

		handleEndOfDigit := func(currentIdx int) {
			digit := line[startIdx:currentIdx]

			xRange := common.NewSlice(startIdx-1, currentIdx, 1)
			yRange := common.NewSlice(lineIdx-1, lineIdx+1, 1)

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

	return strconv.Itoa(common.Sum(products))
}

func (c ChallengeThree) DataFolder() string {
	return "03"
}
