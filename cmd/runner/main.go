package main

import (
	"advent-of-code-2023/internal/challenges"
	"advent-of-code-2023/internal/filereader"
	"fmt"
)

func main() {
	challenge := challenges.NewChallenge04()
	fr := filereader.InputFileReader()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")

	outputOne := challenge.RunPartOne(input)
	outputTwo := challenge.RunPartTwo(input)

	fmt.Printf("Output of part 1 is: %s\n", outputOne)
	fmt.Printf("Output of part 2 is: %s\n", outputTwo)
}
