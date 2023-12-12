package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"strconv"
	"strings"
)

func NewChallenge12() Challenge12 {
	return Challenge12{
		cache: make(map[string]int),
	}
}

const (
	NoSpring = '.'
	Spring   = '#'
	Unknown  = '?'
)

type Challenge12 struct {
	cache map[string]int
}

func (c Challenge12) RunPartOne(input string) string {
	configCounts := 0

	lines := helpers.SplitLines(input)

	for _, line := range lines {
		data := strings.Split(line, " ")
		groups := helpers.StringListOfNumericValuesToSlice(data[1], ",")
		configCounts += c.getOptionCountsForConfig(data[0], groups)
	}

	return strconv.Itoa(configCounts)
}

func (c Challenge12) RunPartTwo(input string) string {
	configCounts := 0

	lines := helpers.SplitLines(input)
	for _, line := range lines {
		data := strings.Split(line, " ")
		groups := helpers.StringListOfNumericValuesToSlice(data[1], ",")
		configCounts += c.getOptionCountsForConfig(repeatConfigs(data[0]), repeatGroups(groups))
	}

	return strconv.Itoa(configCounts)
}

func repeatConfigs(input string) string {
	newString := input
	for i := 0; i < 4; i++ {
		newString += string(Unknown) + input
	}
	return newString
}

func repeatGroups(input []int) []int {
	newSlice := make([]int, 0, len(input)*5)
	for i := 0; i < 5; i++ {
		newSlice = append(newSlice, input...)
	}
	return newSlice
}

func (c Challenge12) cachedValue(key string, value int) int {
	c.cache[key] = value
	return value
}

func (c Challenge12) getOptionCountsForConfig(config string, expectedGroups []int) int {
	cacheKey := config + "-" + helpers.ToString(expectedGroups)

	if value, ok := c.cache[cacheKey]; ok {
		return value
	}

	if len(expectedGroups) == 0 && len(config) == 0 {
		return 1
	}

	if len(expectedGroups) > 0 && len(config) == 0 {
		return 0
	}

	if len(config) == 0 {
		return 0
	}

	expectedSpringsLeft := aoc_math.Sum(expectedGroups)

	if strings.Count(config, string(Spring)) > expectedSpringsLeft {
		return c.cachedValue(cacheKey, 0)
	}

	if strings.Count(config, string(Spring))+strings.Count(config, string(Unknown)) < expectedSpringsLeft {
		return 0
	}

	if config[0] == Unknown {
		value := c.getOptionCountsForConfig(string(Spring)+config[1:], expectedGroups) +
			c.getOptionCountsForConfig(config[1:], expectedGroups)

		return c.cachedValue(cacheKey, value)
	}

	if config[0] == NoSpring {
		value := c.getOptionCountsForConfig(config[1:], expectedGroups)
		return c.cachedValue(cacheKey, value)
	}

	if config[0] == Spring {
		expectedLengthOfGroup := expectedGroups[0]
		if expectedLengthOfGroup > len(config) {
			return c.cachedValue(cacheKey, 0)
		}

		group := config[0:expectedLengthOfGroup]
		if !helpers.AllSatisfies([]rune(group), func(r rune) bool {
			return r == Spring || r == Unknown
		}) {
			return c.cachedValue(cacheKey, 0)
		} else {
			remainingConfig := config[expectedLengthOfGroup:]
			if len(remainingConfig) == 0 {
				return c.getOptionCountsForConfig("", expectedGroups[1:])
			}
			if remainingConfig[0] == Spring {
				return c.cachedValue(cacheKey, 0)
			}

			value := c.getOptionCountsForConfig(config[expectedLengthOfGroup+1:], expectedGroups[1:])
			return c.cachedValue(cacheKey, value)
		}
	}
	panic("this should not happen")
}

func (c Challenge12) DataFolder() string {
	return "12"
}
