package main

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type BinList struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func main() {
	bin := NewBin("123", "egor", true)
	binList := NewBinList("123", "Marta", false)
	fmt.Println(bin, binList)
}
func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func NewBinList(id, name string, private bool) *BinList {
	return &BinList{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
