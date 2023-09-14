package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw04_struct_comparator/book"
)

func main() {
	b1 := book.NewBook("Title1", "Author1", 1928, 500, 4.6)
	fmt.Printf("%v", b1)
}
