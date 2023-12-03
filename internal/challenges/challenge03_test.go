package challenges

import (
	"advent-of-code-2023/internal/filereader"
	"testing"
)

func TestChallengeThree_RunPartOne_ExampleInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallengeThree()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.1.out")

	output := challenge.RunPartOne(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallengeThree_RunPartOne_ChallengeInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallengeThree()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.1.out")

	output := challenge.RunPartOne(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallengeThree_RunPartTwo_ExampleInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallengeThree()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.2.out")

	output := challenge.RunPartTwo(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallengeThree_RunPartTwo_ChallengeInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallengeThree()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.2.out")

	output := challenge.RunPartTwo(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}
