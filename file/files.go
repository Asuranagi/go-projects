package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFiles(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func IsJsonFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}

func WriteFile(content string, name string) {

}
