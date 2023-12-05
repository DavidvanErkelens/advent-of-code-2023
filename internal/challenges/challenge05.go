package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge05() Challenge05 {
	return Challenge05{
		keyOrder: []string{
			"seed-soil",
			"soil-fertilizer",
			"fertilizer-water",
			"water-light",
			"light-temperature",
			"temperature-humidity",
			"humidity-location",
		},
		reversedKeyOrder: []string{
			"humidity-location",
			"temperature-humidity",
			"light-temperature",
			"water-light",
			"fertilizer-water",
			"soil-fertilizer",
			"seed-soil",
		},
	}
}

type Challenge05 struct {
	keyOrder         []string
	reversedKeyOrder []string
}

type Mapping struct {
	FromStart int
	FromEnd   int
	ToStart   int
	ToEnd     int
	Offset    int
}

type SeedRange struct {
	From  int
	Range int
}

func (s SeedRange) InRange(value int) bool {
	return value >= s.From && value < s.From+s.Range
}

func (m Mapping) IsPartOfMapping(value int) bool {
	return value >= m.FromStart && value <= m.FromEnd
}

func (m Mapping) IsPartOfMappingReverse(value int) bool {
	return value >= m.ToStart && value <= m.ToEnd
}

func (m Mapping) Map(value int) int {
	if !m.IsPartOfMapping(value) {
		return 0
	}

	return value + m.Offset
}

func (m Mapping) MapReverse(value int) int {
	if !m.IsPartOfMappingReverse(value) {
		return 0
	}

	return value - m.Offset
}

func (c Challenge05) parseMappings(mapDefinitions []string) map[string][]Mapping {
	maps := make(map[string][]Mapping)

	for _, definition := range mapDefinitions {
		lines := helpers.SplitLines(definition)
		headerData := strings.Split(lines[0], " ")
		fromToData := strings.Split(headerData[0], "-")
		from := fromToData[0]
		to := fromToData[2]
		key := fmt.Sprintf("%s-%s", from, to)

		mappingsForKey := make([]Mapping, 0)
		for _, line := range lines[1:] {
			numbers := helpers.StringListOfNumericValuesToSlice(line, " ")
			mappingsForKey = append(mappingsForKey, Mapping{
				FromStart: numbers[1],
				FromEnd:   numbers[1] + numbers[2] - 1,
				ToStart:   numbers[0],
				ToEnd:     numbers[0] + numbers[2] - 1,
				Offset:    numbers[0] - numbers[1],
			})
		}
		slices.SortFunc(mappingsForKey, func(a, b Mapping) int {
			return a.ToStart - b.ToStart
		})

		maps[key] = mappingsForKey
	}

	return maps
}

func (c Challenge05) getLocationValue(mappings map[string][]Mapping, input int) int {
	currentValue := input
	for _, key := range c.keyOrder {
		mappingsForKey := mappings[key]
		for _, mapping := range mappingsForKey {
			if mapping.IsPartOfMapping(currentValue) {
				currentValue = mapping.Map(currentValue)
				break
			}
		}
	}

	return currentValue
}

func (c Challenge05) getSeedNumberForLocation(mappings map[string][]Mapping, input int) int {
	currentValue := input
	for _, key := range c.reversedKeyOrder {
		mappingsForKey := mappings[key]
		for _, mapping := range mappingsForKey {
			if mapping.IsPartOfMappingReverse(currentValue) {
				currentValue = mapping.MapReverse(currentValue)
				break
			}
		}
	}

	return currentValue
}

func (c Challenge05) getSoilNumbersForSeeds(mappings map[string][]Mapping, seeds []int) []int {
	locations := make([]int, 0)

	for _, seed := range seeds {
		soil := c.getLocationValue(mappings, seed)
		locations = append(locations, soil)
	}

	return locations
}

func (c Challenge05) RunPartOne(input string) string {
	paragraphs := helpers.SplitParagraphs(input)
	mapDefinitions := paragraphs[1:]

	mappings := c.parseMappings(mapDefinitions)

	seedLineData := strings.Split(paragraphs[0], ":")
	seeds := helpers.StringListOfNumericValuesToSlice(strings.TrimSpace(seedLineData[1]), " ")

	locations := c.getSoilNumbersForSeeds(mappings, seeds)
	minimumSoil := slices.Min(locations)

	return strconv.Itoa(minimumSoil)
}

// Not the fastest solution (takes around 5s) but good enough ¯\_(ツ)_/¯
func (c Challenge05) RunPartTwo(input string) string {
	paragraphs := helpers.SplitParagraphs(input)
	mapDefinitions := paragraphs[1:]

	mappings := c.parseMappings(mapDefinitions)

	seedLineData := strings.Split(paragraphs[0], ":")
	seedData := helpers.StringListOfNumericValuesToSlice(strings.TrimSpace(seedLineData[1]), " ")
	seeds := make([]SeedRange, 0)

	for i := 0; i < len(seedData); i += 2 {
		seeds = append(seeds, SeedRange{
			From:  seedData[i],
			Range: seedData[i+1],
		})
	}

	result := 0
	i := 1
	for result == 0 {
		startSeed := c.getSeedNumberForLocation(mappings, i)

		for _, seed := range seeds {
			if seed.InRange(startSeed) {
				result = i
				break
			}
		}
		i += 1
	}

	return strconv.Itoa(result)
}

func (c Challenge05) DataFolder() string {
	return "05"
}
