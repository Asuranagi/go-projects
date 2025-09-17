package main

import (
	"true-Pro/file"
	"true-Pro/storage"
)

func main() {
	file.ReadFiles("file.json")
	file.IsJsonFile("file.json")
	storage.ReadBinListFromJson("file.json")
}
