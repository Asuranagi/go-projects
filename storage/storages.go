package storage

import (
	"encoding/json"

	"os"
	"true-Pro/bins"
)

type StorageService interface {
	SaveBinToJSON(bin *bins.Bin, filename string) error
	ReadBinListFromJson(filename string) ([]bins.Bin, error)
}

type StorageServiceImpl struct{}

func NewStorageService() StorageService {
	return &StorageServiceImpl{}
}

func (s *StorageServiceImpl) SaveBinToJSON(bin *bins.Bin, filename string) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (s *StorageServiceImpl) ReadBinListFromJson(filename string) ([]bins.Bin, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var binList []bins.Bin
	err = json.Unmarshal(data, &binList)
	return binList, err
}
