package main

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw05_shapes/internal/shape"
)

func calculateArea(s any) (float32, error) {
	if v, ok := s.(shape.Shape); ok {
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

	shapes = append(shapes, r, c, t, i)

	for _, shape := range shapes {
		if area, err := calculateArea(shape); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(shape)
			fmt.Println("Area:", area)
		}
	}
}
