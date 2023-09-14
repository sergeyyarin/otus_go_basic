package book

func newId() func() int {
	var id int = 0
	return func() int {
		id += 1
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
		id:     newId()(),
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}
