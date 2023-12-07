package challenges

import (
	"testing"
)

func TestChallenge07_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallenge07()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallenge07_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallenge07()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallenge07_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallenge07()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallenge07_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallenge07()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
