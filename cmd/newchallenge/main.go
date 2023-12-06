package main

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/projectpath"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	challengeNumber := fmt.Sprintf("%02d", getNextChallengeNumber())
	createInputDirectory(challengeNumber)
	createChallenge(challengeNumber)
}

func getNextChallengeNumber() int {
	files, err := os.ReadDir(projectpath.Root + "/input")
	if err != nil {
		log.Fatal(err)
	}

	lastChallenge := 0
	for _, file := range files {
		if file.IsDir() {
			lastChallenge = max(lastChallenge, helpers.NaiveStringToInt(file.Name()))
		}
	}
	return lastChallenge + 1
}

func createInputDirectory(challenge string) {
	inputPath := projectpath.Root + "/input/" + challenge
	_ = os.Mkdir(inputPath, os.ModePerm)

	filesToCreate := []string{
		"example.in", "example.1.out", "example.2.out", "challenge.in", "challenge.1.out", "challenge.2.out",
	}

	for _, file := range filesToCreate {
		_, _ = os.Create(inputPath + "/" + file)
	}
}

func createChallenge(challenge string) {
	challengeTemplateFile, _ := os.ReadFile(projectpath.Root + "/internal/challenges/challenge.tpl")
	testTemplateFile, _ := os.ReadFile(projectpath.Root + "/internal/challenges/challenge_test.tpl")
	challengeTemplate := string(challengeTemplateFile)
	testTemplate := string(testTemplateFile)

	challengeTemplate = strings.ReplaceAll(challengeTemplate, "CHALLENGE_NUMBER", challenge)
	testTemplate = strings.ReplaceAll(testTemplate, "CHALLENGE_NUMBER", challenge)

	os.WriteFile(projectpath.Root+"/internal/challenges/challenge"+challenge+".go", []byte(challengeTemplate), os.ModePerm)
	os.WriteFile(projectpath.Root+"/internal/challenges/challenge"+challenge+"_test.go", []byte(testTemplate), os.ModePerm)
}
