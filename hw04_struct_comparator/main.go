package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw04_struct_comparator/book"
)

func compareBooks(lhs *book.Book, rhs *book.Book, mode book.CompareMode) bool {
	cmp := book.NewComparator(mode, lhs, rhs)
	res, err := cmp.Compare()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func main() {
	a := book.NewBook("Title1", "Author1", 1901, 100, 8.7)
	b := book.NewBook("Title2", "Author2", 2009, 1450, 4.8)

	if compareBooks(&a, &b, book.Size) {
		fmt.Printf("%s by %s is bigger\n", a.GetTitle(), a.GetAuthor())
	} else {
		fmt.Printf("%s by %s is bigger\n", b.GetTitle(), b.GetAuthor())
	}

	if compareBooks(&a, &b, book.Year) {
		fmt.Printf("%s by %s is more recent\n", a.GetTitle(), a.GetAuthor())
	} else {
		fmt.Printf("%s by %s is more recent\n", b.GetTitle(), b.GetAuthor())
	}

	if compareBooks(&a, &b, book.Rate) {
		fmt.Printf("%s by %s is more popular\n", a.GetTitle(), a.GetAuthor())
	} else {
		fmt.Printf("%s by %s is more popular\n", b.GetTitle(), b.GetAuthor())
	}
}
