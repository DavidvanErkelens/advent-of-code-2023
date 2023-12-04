package helpers

import (
	"strconv"
	"strings"
)

func SplitChunks(input string, length int) []string {
	chunks := make([]string, 0)

	for i := 0; i < len(input); i += length {
		substr := input[i : i+length]
		chunks = append(chunks, substr)
	}

	return chunks
}

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func NaiveStringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}

func StringListOfNumericValuesToSlice(s string, separator string) []int {
	splitValues := strings.Split(s, separator)
	filtered := Filter(splitValues, func(val string) bool {
		return strings.TrimSpace(val) != ""
	})
	return Map(filtered, NaiveStringToInt)
}
