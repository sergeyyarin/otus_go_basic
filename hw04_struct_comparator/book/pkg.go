package book

import "errors"

var id int

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func newID() int {
	id++
	return id
}

func NewBook(title string, author string, year int, size int, rate float32) Book {
	return Book{
		id:     newID(),
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) SetID(newID int) {
	b.id = newID
}

func (b *Book) SetTitle(newTitle string) {
	b.title = newTitle
}

func (b *Book) SetAuthor(newAuthor string) {
	b.author = newAuthor
}

func (b *Book) SetYear(newYear int) {
	b.year = newYear
}

func (b *Book) SetSize(newSize int) {
	b.size = newSize
}

func (b *Book) SetRate(newRate float32) {
	b.rate = newRate
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float32 {
	return b.rate
}

type CompareMode int

const (
	Year CompareMode = iota
	Size
	Rate
)

type Comparator struct {
	Mode CompareMode
	LHS  *Book
	RHS  *Book
}

func NewComparator(mode CompareMode, lhs *Book, rhs *Book) *Comparator {
	return &Comparator{
		Mode: mode,
		LHS:  lhs,
		RHS:  rhs,
	}
}

func (c *Comparator) Compare() (bool, error) {
	var res bool
	var err error
	switch c.Mode {
	case Year:
		if c.LHS.GetYear() > c.RHS.GetYear() {
			res = true
		}
	case Size:
		if c.LHS.GetSize() > c.RHS.GetSize() {
			res = true
		}
	case Rate:
		if c.LHS.GetRate() > c.RHS.GetRate() {
			res = true
		}
	default:
		err = errors.New("unknown comparison mode")
	}
	return res, err
}
