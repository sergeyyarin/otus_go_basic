package shape

import "fmt"

type Triangle struct {
	base   float32
	height float32
}

func NewTriangle(b float32, h float32) *Triangle {
	return &Triangle{
		base:   b,
		height: h,
	}
}

func (t *Triangle) SetBase(val float32) {
	t.base = val
}

func (t *Triangle) SetHeight(val float32) {
	t.height = val
}

func (t Triangle) Base() float32 {
	return t.base
}

func (t Triangle) Height() float32 {
	return t.height
}

func (t Triangle) Area() float32 {
	return 0.5 * (t.base * t.height)
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle: base %.f, height %.f", t.base, t.height)
}
