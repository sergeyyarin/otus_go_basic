package shape

type Rectangle struct {
	length float32
	width  float32
}

func (r *Rectangle) SetLength(val float32) {
	r.length = val
}

func (r *Rectangle) SetWidth(val float32) {
	r.width = val
}

func (r Rectangle) Length() float32 {
	return r.length
}

func (r Rectangle) Width() float32 {
	return r.width
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}
