package file

import (
	"os"
	"path/filepath"
	"strings"
)

type FileService interface {
	ReadFiles(name string) ([]byte, error)
	IsJsonFile(filename string) bool
	WriteFile(content string, name string)
}
type FileServiceImpl struct{}

func NewFileService() FileService {
	return &FileServiceImpl{}
}
func (f *FileServiceImpl) ReadFiles(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FileServiceImpl) IsJsonFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}
func (f *FileServiceImpl) WriteFile(content string, name string) {

}
