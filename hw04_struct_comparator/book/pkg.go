package book

import "errors"

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

func (c *Comparator) Compare() (*Book, error) {
	var res *Book
	var err error
	switch c.Mode {
	case Year:
		if c.LHS.GetYear() > c.RHS.GetYear() {
			res = c.LHS
		} else {
			res = c.RHS
		}
	case Size:
		if c.LHS.GetSize() > c.RHS.GetSize() {
			res = c.LHS
		} else {
			res = c.RHS
		}
	case Rate:
		if c.LHS.GetRate() > c.RHS.GetRate() {
			res = c.LHS
		} else {
			res = c.RHS
		}
	default:
		res = nil
		err = errors.New("unknown comparison mode")
	}
	return res, err
}
