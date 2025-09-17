package file

import (
	"os"
)

func ReadFiles(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func WriteFile(content string, name string) {

}
