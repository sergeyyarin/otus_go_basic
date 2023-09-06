package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sergeyyarin/otus_go_basic/hw03_chessboard/board"
)

func main() {
	var size int

	fmt.Print("Enter the size of your board: ")

	if _, err := fmt.Scanf("%d", &size); err != nil {
		fmt.Println("can't read board size: ", err)
		os.Exit(1)
	}

	if size <= 0 {
		fmt.Println(errors.New("size is not valid"))
		os.Exit(1)
	}

	if size%2 != 0 {
		fmt.Println(errors.New("size is not even"))
		os.Exit(1)
	}

	board.Build(size)
}
