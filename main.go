package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}
type BinList struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {
	bin := NewBin("123", "egor", true)
	binList := NewBinList("123", "Marta", false)
	fmt.Println(bin, binList)
}
func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
func NewBinList(id, name string, private bool) *BinList {
	return &BinList{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
