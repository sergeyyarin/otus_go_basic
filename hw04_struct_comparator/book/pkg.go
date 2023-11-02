package book

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

func (c *BookComparator) Compare() *Book {
	switch c.Mode {
	case Year:
		if c.Lhs.year > c.Rhs.year {
			return c.Lhs
		} else {
			return c.Rhs
		}
	case Size:
		if c.Lhs.size > c.Rhs.size {
			return c.Lhs
		} else {
			return c.Rhs
		}
	case Rate:
		if c.Lhs.rate > c.Rhs.rate {
			return c.Lhs
		} else {
			return c.Rhs
		}
	}
	return nil
}
