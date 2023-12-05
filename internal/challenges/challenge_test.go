package challenges

import (
	"advent-of-code-2023/internal/filereader"
	"testing"
)

func RunChallengeTest(t *testing.T, dataDir string, inputFile string, outputFile string, function func(string) string) {
	fr := filereader.InputFileReader()

	if !fr.FileExists(dataDir, inputFile) {
		t.Skip("Input file does not exist")
	}

	input := fr.GetFileDataOrPanic(dataDir, inputFile)
	expectedOutput := fr.GetFileDataOrPanic(dataDir, outputFile)

	output := function(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: got '%s', expected '%s'", output, expectedOutput)
	}
}
