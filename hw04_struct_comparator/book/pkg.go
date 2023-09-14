package book

func newID() func() int {
	var id int
	return func() int {
		id++
		return id
	}
}

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(title, author string, year, size int, rate float64) Book {
	return Book{
		id:     newID()(),
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}
