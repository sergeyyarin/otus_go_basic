package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw04_struct_comparator/book"
)

func compare(lhs *book.Book, rhs *book.Book, mode book.CompareMode) bool {
	cmp := book.NewComparator(mode)
	return cmp.Compare(lhs, rhs)
}

func main() {
	a := book.NewBook("Title1", "Author1", 1901, 100, 8.7)
	b := book.NewBook("Title2", "Author2", 2009, 1450, 4.8)

	if compare(&a, &b, book.Size) {
		fmt.Printf("%s by %s is bigger\n", a.Title(), a.Author())
	} else {
		fmt.Printf("%s by %s is bigger\n", b.Title(), b.Author())
	}

	if compare(&a, &b, book.Year) {
		fmt.Printf("%s by %s is more recent\n", a.Title(), a.Author())
	} else {
		fmt.Printf("%s by %s is more recent\n", b.Title(), b.Author())
	}

	if compare(&a, &b, book.Rate) {
		fmt.Printf("%s by %s is more popular\n", a.Title(), a.Author())
	} else {
		fmt.Printf("%s by %s is more popular\n", b.Title(), b.Author())
	}
}
