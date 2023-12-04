package challenges

import (
	"advent-of-code-2023/internal/filereader"
	"testing"
)

func TestChallenge03_RunPartOne_ExampleInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallenge03()

	if !fr.FileExists(challenge.DataFolder(), "example.in") {
		t.Skip("Input file does not exist")
	}

	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.1.out")

	output := challenge.RunPartOne(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallenge03_RunPartOne_ChallengeInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallenge03()

	if !fr.FileExists(challenge.DataFolder(), "challenge.in") {
		t.Skip("Input file does not exist")
	}

	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.1.out")

	output := challenge.RunPartOne(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallenge03_RunPartTwo_ExampleInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallenge03()

	if !fr.FileExists(challenge.DataFolder(), "example.in") {
		t.Skip("Input file does not exist")
	}

	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "example.2.out")

	output := challenge.RunPartTwo(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}

func TestChallenge03_RunPartTwo_ChallengeInput(t *testing.T) {
	fr := filereader.InputFileReader()
	challenge := NewChallenge03()

	if !fr.FileExists(challenge.DataFolder(), "challenge.in") {
		t.Skip("Input file does not exist")
	}

	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")
	expectedOutput := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.2.out")

	output := challenge.RunPartTwo(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}
