package main

import (
	"advent-of-code-2023/internal/challenges"
	"advent-of-code-2023/internal/filereader"
	"fmt"
)

func main() {
	challenge := challenges.NewChallengeTwo()
	fr := filereader.InputFileReader()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "nien.in")

	output := challenge.RunPartOne(input)

	fmt.Printf("Output is: %s\n", output)
}
