package main

import (
	"fmt"
)

func calculateArea(s any) (float32, error) {
	if v, ok := s.(Circle); ok {
		fmt.Printf("Circe: radius %f\n", v.Radius())
		return v.Area(), nil
	}
	if v, ok := s.(Rectangle); ok {
		fmt.Printf("Rectangle: length %f, width %f\n", v.Length(), v.Width())
		return v.Area(), nil
	}
	if v, ok := s.(Triangle); ok {
		fmt.Printf("Triangle: base %f, height %f\n", v.Base(), v.Height())
		return v.Area(), nil
	}
	return 0, fmt.Errorf("interface Shape is not implemented for the type")
}

func main() {
	var shapes []any

	var c Circle
	c.SetRadius(5)

	var r Rectangle
	r.SetLength(10)
	r.SetWidth(5)

	var t Triangle
	t.SetBase(8)
	t.SetHeight(6)

	var i int

	shapes = append(shapes, i, c, r, t)

	for shape := range shapes {
		if area, err := calculateArea(shape); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Area: ", area)
		}
	}
}
