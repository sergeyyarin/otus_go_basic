package book

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

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

type CompareMode int

const (
	Year CompareMode = iota
	Size
	Rate
)

type (
	CompareMethodInt     func(int, int) bool
	CompareMethodFloat32 func(float32, float32) bool
)

type Comparator struct {
	Mode         CompareMode
	CompareInt   CompareMethodInt
	CompareFloat CompareMethodFloat32
}

func NewComparator(mode CompareMode) *Comparator {
	return &Comparator{
		Mode:         mode,
		CompareInt:   func(i1, i2 int) bool { return i1 > i2 },
		CompareFloat: func(f1, f2 float32) bool { return f1 > f2 },
	}
}

func (c *Comparator) Compare(lhs *Book, rhs *Book) bool {
	var res bool
	switch c.Mode {
	case Year:
		res = c.CompareInt(lhs.Year(), rhs.Year())
	case Size:
		res = c.CompareInt(lhs.Size(), rhs.Size())
	case Rate:
		res = c.CompareFloat(lhs.Rate(), rhs.Rate())
	}
	return res
}
