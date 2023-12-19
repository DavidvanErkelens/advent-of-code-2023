package challenges

import (
	"advent-of-code-2023/internal/helpers"
	math2 "advent-of-code-2023/internal/helpers/aoc_math"
	"advent-of-code-2023/internal/helpers/aoc_range"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func NewChallenge01() Challenge01 {
	return Challenge01{}
}

type Challenge01 struct {
}

func (c Challenge01) RunPartOne(input string) string {
	lines := helpers.SplitLines(input)
	numbers := make([]int, 0)

	for _, line := range lines {
		chars := []rune(line)
		digits := helpers.Filter(chars, unicode.IsDigit)
		num := []rune{digits[0], digits[len(digits)-1]}
		value, _ := strconv.Atoi(string(num))

		numbers = append(numbers, value)
	}

	return strconv.Itoa(math2.Sum(numbers))
}

func (c Challenge01) RunPartTwo(input string) string {
	lines := helpers.SplitLines(input)
	numbers := make([]int, 0)

	values := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, val := range aoc_range.NewRangeSlice(1, 9) {
		values[strconv.Itoa(val)] = val
	}

	getFirstNumber := func(line string, isReverse bool) int {
		minIdx := math.MaxInt
		minIdxValue := 0

		for search, value := range values {
			if isReverse {
				search = helpers.Reverse(search)
			}
			idx := strings.Index(line, search)
			if idx >= 0 && idx < minIdx {
				minIdx = idx
				minIdxValue = value
			}
		}

		return minIdxValue
	}

	for _, line := range lines {
		firstNum := getFirstNumber(line, false)
		lastNum := getFirstNumber(helpers.Reverse(line), true)
		value := firstNum*10 + lastNum

		numbers = append(numbers, value)
	}

	return strconv.Itoa(math2.Sum(numbers))
}

func (c Challenge01) DataFolder() string {
	return "01"
}
