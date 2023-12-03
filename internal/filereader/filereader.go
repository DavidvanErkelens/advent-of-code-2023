package filereader

import (
	"advent-of-code-2023/internal/projectpath"
	"os"
	"strings"
)

func InputFileReader() FileReader {
	return FileReader{basePath: projectpath.Root + "/input/data/"}
}

type FileReader struct {
	basePath string
}

func (reader *FileReader) GetInputFilesForChallenge(dataFolder string) ([]string, error) {
	entries, err := os.ReadDir(reader.basePath + dataFolder)
	if err != nil {
		return nil, err
	}

	inputFiles := make([]string, 0)

	for _, e := range entries {
		fileName := e.Name()
		if strings.HasSuffix(fileName, ".in") {
			inputFiles = append(inputFiles, fileName)
		}
	}

	return inputFiles, nil
}

func (reader *FileReader) GetFileData(dataFolder string, fileName string) (string, error) {
	path := reader.getPath(dataFolder, fileName)
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (reader *FileReader) getPath(dataFolder string, fileName string) string {
	return reader.basePath + dataFolder + "/" + fileName
}

//func (reader *FileReader) GetFileDataForChallengeOrPanic(challenge *challenges.Challenge, fileName string) string {
//	return reader.GetFileDataOrPanic(challenge.DataFolder(), fileName)
//}

func (reader *FileReader) GetFileDataOrPanic(dataFolder string, fileName string) string {
	data, err := reader.GetFileData(dataFolder, fileName)
	if err != nil {
		panic(err)
	}
	return data
}
