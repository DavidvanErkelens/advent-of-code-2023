package challenges

import (
	"advent-of-code-2023/internal/common"
	"math"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge04() Challenge04 {
	return Challenge04{}
}

type Challenge04 struct {
}

func (c Challenge04) RunPartOne(input string) string {
	lines := common.SplitLines(input)
	scoreSums := 0

	for _, line := range lines {
		configDataSplit := strings.Split(line, ":")
		numbersSplit := strings.Split(configDataSplit[1], "|")

		cardNumbersStrings := strings.Split(
			strings.TrimSpace(numbersSplit[0]),
			" ",
		)
		cardNumbers := common.Map(cardNumbersStrings, func(input string) int {
			parsed, _ := strconv.Atoi(input)
			return parsed
		})
		cardNumbers = common.Filter(cardNumbers, func(val int) bool {
			return val > 0
		})

		winningNumbersStrings := strings.Split(
			strings.TrimSpace(numbersSplit[1]),
			" ",
		)
		winningNumbers := common.Map(winningNumbersStrings, func(input string) int {
			parsed, _ := strconv.Atoi(input)
			return parsed
		})

		winners := 0

		for _, cardNumber := range cardNumbers {
			if slices.Contains(winningNumbers, cardNumber) {
				winners += 1
			}
		}

		if winners > 0 {
			outcome := int(math.Pow(2, float64(winners-1)))
			scoreSums += outcome
		}
	}

	return strconv.Itoa(scoreSums)
}

func (c Challenge04) RunPartTwo(input string) string {
	lines := common.SplitLines(input)
	numberOfCards := make(map[int]int, len(lines)+1)

	for _, idx := range common.NewSlice(1, len(lines), 1) {
		numberOfCards[idx] = 1
	}

	for idx, line := range lines {
		configDataSplit := strings.Split(line, ":")
		numbersSplit := strings.Split(configDataSplit[1], "|")

		cardNumbersStrings := strings.Split(
			strings.TrimSpace(numbersSplit[0]),
			" ",
		)
		cardNumbers := common.Map(cardNumbersStrings, func(input string) int {
			parsed, _ := strconv.Atoi(input)
			return parsed
		})
		cardNumbers = common.Filter(cardNumbers, func(val int) bool {
			return val > 0
		})

		winningNumbersStrings := strings.Split(
			strings.TrimSpace(numbersSplit[1]),
			" ",
		)
		winningNumbers := common.Map(winningNumbersStrings, func(input string) int {
			parsed, _ := strconv.Atoi(input)
			return parsed
		})

		winners := 0

		for _, cardNumber := range cardNumbers {
			if slices.Contains(winningNumbers, cardNumber) {
				winners += 1
			}
		}

		if winners == 0 {
			continue
		}

		cardNumber := idx + 1
		cardsForThisNumber := numberOfCards[cardNumber]

		cardsToAdd := common.NewSlice(cardNumber+1, cardNumber+winners, 1)

		for _, extraCardIdx := range cardsToAdd {
			numberOfCards[extraCardIdx] += cardsForThisNumber
		}
	}

	totalCards := 0
	for _, cards := range numberOfCards {
		totalCards += cards
	}

	return strconv.Itoa(totalCards)
}

func (c Challenge04) DataFolder() string {
	return "04"
}
