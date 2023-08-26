package main

import (
	"errors"
	"fmt"
)

func buildRow(length int) (row string, err error) {
	if valid := length > 0; !valid {
		err = errors.New("row len is not valid")
		return row, err
	}
	if even := length%2 == 0; !even {
		err = errors.New("row len is not even")
		return row, err
	}

	const (
		space = " "
		hash  = "#"
	)

	for i := 0; i <= length; i++ {
		if i%2 == 0 {
			row += hash
		} else {
			row += space
		}
	}

	return row, err
}

func main() {
	length := 8
	row, err := buildRow(length)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Board: ", row)
	}
}
