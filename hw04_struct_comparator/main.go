package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw04_struct_comparator/book"
)

func compareBooks(lhs *book.Book, rhs *book.Book, mode book.CompareMode) *book.Book {
	compare := book.BookComparator{Mode: mode, Lhs: lhs, Rhs: rhs}
	if res := compare.Compare(); res != nil {
		return res
	}
	return nil
}

func main() {
	a := book.Book{}
	b := book.Book{}

	a.SetID(1)
	a.SetTitle("Title1")
	a.SetAuthor("Author1")
	a.SetYear(1901)
	a.SetSize(100)
	a.SetRate(8.7)

	b.SetID(2)
	b.SetTitle("Title2")
	b.SetAuthor("Author2")
	b.SetYear(2009)
	b.SetSize(1450)
	b.SetRate(4.8)

	bigger := compareBooks(&a, &b, book.Size)
	if bigger != nil {
		fmt.Printf("%s by %s is bigger\n", bigger.GetTitle(), bigger.GetAuthor())
	}

	newer := compareBooks(&a, &b, book.Year)
	if newer != nil {
		fmt.Printf("%s by %s is more recent\n", newer.GetTitle(), newer.GetAuthor())
	}

	best := compareBooks(&a, &b, book.Rate)
	if best != nil {
		fmt.Printf("%s by %s is more popular\n", best.GetTitle(), best.GetAuthor())
	}
}
