package main

import (
	"github.com/sergeyyarin/otus_go_basic/hw02_fix_app/printer"
	"github.com/sergeyyarin/otus_go_basic/hw02_fix_app/reader"
	"github.com/sergeyyarin/otus_go_basic/hw02_fix_app/types"
	"fmt"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)

	if err != nil {
		fmt.Println(err)
	} else {
		printer.PrintStaff(staff)
	}
}
