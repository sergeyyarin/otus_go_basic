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

type CompareMethodInt func(int, int) bool
type CompareMethodFloat32 func(float32, float32) bool

type Comparator struct {
	Mode         CompareMode
	CompareInt   CompareMethodInt
	compareFloat CompareMethodFloat32
}

func NewComparator(mode CompareMode) *Comparator {
	return &Comparator{
		Mode:         mode,
		CompareInt:   func(i1, i2 int) bool { return i1 > i2 },
		compareFloat: func(f1, f2 float32) bool { return f1 > f2 },
	}
}

func (c *Comparator) Compare(lhs *Book, rhs *Book) bool {
	var res bool
	switch c.Mode {
	case Year:
		res = c.CompareInt(lhs.GetYear(), rhs.GetYear())
	case Size:
		res = c.CompareInt(lhs.GetSize(), rhs.GetSize())
	case Rate:
		res = c.compareFloat(lhs.GetRate(), rhs.GetRate())
	}
	return res
}
