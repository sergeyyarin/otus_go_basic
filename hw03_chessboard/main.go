package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sergeyyarin/otus_go_basic/hw03_chessboard/board"
)

func main() {
	size := 8

	if valid := size > 0; !valid {
		fmt.Println(errors.New("size is not valid"))
		os.Exit(1)
	}

	if even := size%2 == 0; !even {
		fmt.Println(errors.New("size is not even"))
		os.Exit(1)
	}

	board.Build(size)
}
