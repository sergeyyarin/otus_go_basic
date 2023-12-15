package shape

import "fmt"

const (
	Pi = 3.1415
)

type Circle struct {
	radius float32
}

func NewCircle(r float32) *Circle {
	return &Circle{
		radius: r,
	}
}

func (c *Circle) SetRadius(val float32) {
	c.radius = val
}

func (c Circle) Radius() float32 {
	return c.radius
}

func (c Circle) Area() float32 {
	return Pi * (c.radius * c.radius)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius %.f", c.radius)
}
