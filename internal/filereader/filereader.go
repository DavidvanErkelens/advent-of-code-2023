package filereader

import (
	"advent-of-code-2023/internal/projectpath"
	"os"
)

func InputFileReader() FileReader {
	return FileReader{basePath: projectpath.Root + "/input/data/"}
}

type FileReader struct {
	basePath string
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

func (reader *FileReader) GetFileDataOrPanic(dataFolder string, fileName string) string {
	data, err := reader.GetFileData(dataFolder, fileName)
	if err != nil {
		panic(err)
	}
	return data
}

func (reader *FileReader) FileExists(dataFolder string, fileName string) bool {
	path := reader.getPath(dataFolder, fileName)
	_, err := os.ReadFile(path)
	return !os.IsNotExist(err)
}
