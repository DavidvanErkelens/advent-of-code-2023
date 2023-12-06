package challenges

import (
	"testing"
)

func TestChallenge01_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallenge01()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallenge01_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallenge01()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallenge01_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallenge01()
	RunChallengeTest(t, challenge.DataFolder(), "example.2.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallenge01_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallenge01()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
