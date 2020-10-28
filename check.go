package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	f, err := xlsx.OpenFile("genesis_file.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(f)
}
