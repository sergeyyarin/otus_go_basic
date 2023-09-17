package book

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

func NewBook() Book {
	return Book{
		id: newID()(),
	}
}

func (r Book) GetID() int {
	return r.id
}

func (r Book) GetTitle() string {
	return r.title
}

func (r Book) GetAuthor() string {
	return r.author
}

func (r Book) GetYear() int {
	return r.year
}

func (r Book) GetSize() int {
	return r.size
}

func (r Book) GetRate() float64 {
	return r.rate
}

func (r *Book) SetTitle(new string) {
	r.title = new
}

func (r *Book) SetAuthor(new string) {
	r.author = new
}

func (r *Book) SetYear(new int) {
	r.year = new
}

func (r *Book) SetSize(new int) {
	r.size = new
}

func (r *Book) SetRate(new float64) {
	r.rate = new
}
