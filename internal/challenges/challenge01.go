package challenges

import (
	"advent-of-code-2023/internal/common"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func NewChallengeOne() ChallengeOne {
	return ChallengeOne{}
}

type ChallengeOne struct {
}

func (c ChallengeOne) RunPartOne(input string) string {
	lines := common.SplitLines(input)
	numbers := make([]int, 0)

	for _, line := range lines {
		chars := []rune(line)
		digits := common.Filter(chars, unicode.IsDigit)
		num := []rune{digits[0], digits[len(digits)-1]}
		value, _ := strconv.Atoi(string(num))

		numbers = append(numbers, value)
	}

	return strconv.Itoa(common.Sum(numbers))
}

func (c ChallengeOne) RunPartTwo(input string) string {
	lines := common.SplitLines(input)
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

	for _, val := range common.NewSlice(1, 9, 1) {
		values[strconv.Itoa(val)] = val
	}

	getFirstNumber := func(line string, isReverse bool) int {
		minIdx := math.MaxInt
		minIdxValue := 0

		for search, value := range values {
			if isReverse {
				search = common.Reverse(search)
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
		lastNum := getFirstNumber(common.Reverse(line), true)
		value := firstNum*10 + lastNum

		numbers = append(numbers, value)
	}

	return strconv.Itoa(common.Sum(numbers))
}

func (c ChallengeOne) DataFolder() string {
	return "01"
}
