package main

import (
	"true-Pro/file"
	"true-Pro/storage"
)

func main() {
	FileService := file.NewFileService()
	StorageService := storage.NewStorageService()
	FileService.ReadFiles("file.json")
	FileService.IsJsonFile("file.json")
	StorageService.ReadBinListFromJson("file.json")
}
