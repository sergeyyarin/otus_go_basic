package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw04_struct_comparator/book"
)

func main() {
	b1 := book.NewBook()
	b1.SetAuthor("Author1")
	b1.SetTitle("Title1")
	b1.SetYear(1928)
	b1.SetSize(500)
	b1.SetRate(4.6)

	b2 := book.NewBook()
	b2.SetAuthor("Author2")
	b2.SetTitle("Title2")
	b2.SetYear(1828)
	b2.SetSize(300)
	b2.SetRate(2.6)

	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", b2)
}
