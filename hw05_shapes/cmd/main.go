package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/tree/hw05_shapes/hw05_shapes/internal/shape"
)

func calculateArea(s any) (float32, error) {
	if v, ok := s.(shape.Circle); ok {
		fmt.Printf("Circe: radius %f\n", v.Radius())
		return v.Area(), nil
	}
	if v, ok := s.(shape.Rectangle); ok {
		fmt.Printf("Rectangle: length %f, width %f\n", v.Length(), v.Width())
		return v.Area(), nil
	}
	if v, ok := s.(shape.Triangle); ok {
		fmt.Printf("Triangle: base %f, height %f\n", v.Base(), v.Height())
		return v.Area(), nil
	}
	return 0, fmt.Errorf("interface Shape is not implemented for the type")
}

func main() {
	var shapes []any

	var c shape.Circle
	c.SetRadius(5)

	var r shape.Rectangle
	r.SetLength(10)
	r.SetWidth(5)

	var t shape.Triangle
	t.SetBase(8)
	t.SetHeight(6)

	var i int

	shapes = append(shapes, c, r, t, i)

	for _, shape := range shapes {
		if area, err := calculateArea(shape); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Area: ", area)
		}
	}
}
