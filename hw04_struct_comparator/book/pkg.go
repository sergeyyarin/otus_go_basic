package book

import "errors"

var id int

func newID() func() int {
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

func (b *Book) SetRate(newRate float64) {
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

func (b *Book) GetRate() float64 {
	return b.rate
}

type CompareMode int

const (
	Year CompareMode = iota
	Size
	Rate
)

type BookComparator struct {
	Mode CompareMode
	Lhs  *Book
	Rhs  *Book
}

func NewBookComparator(mode CompareMode, lhs *Book, rhs *Book) *BookComparator {
	return &BookComparator{
		Mode: mode,
		Lhs:  lhs,
		Rhs:  rhs,
	}
}

func (c *BookComparator) Compare() (*Book, error) {
	var res *Book
	var err error = nil
	switch c.Mode {
	case Year:
		if c.Lhs.GetYear() > c.Rhs.GetYear() {
			res = c.Lhs
		} else {
			res = c.Rhs
		}
	case Size:
		if c.Lhs.GetSize() > c.Rhs.GetSize() {
			res = c.Lhs
		} else {
			res = c.Rhs
		}
	case Rate:
		if c.Lhs.GetRate() > c.Rhs.GetRate() {
			res = c.Lhs
		} else {
			res = c.Rhs
		}
	default:
		res = nil
		err = errors.New("unknown comparison mode")
	}
	return res, err
}
