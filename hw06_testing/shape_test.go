package hw06testing

import (
	"testing"

	"github.com/sergeyyarin/otus_go_basic/hw06_testing/shape"
	"github.com/stretchr/testify/require"
)

func TestShapeInterface(t *testing.T) {
	testCases := []struct {
		desc   string
		figure shape.Shape
	}{
		{
			desc:   "Circle",
			figure: shape.NewCircle(5),
		},
		{
			desc:   "Triangle",
			figure: shape.NewTriangle(8, 6),
		},
		{
			desc:   "Rectangle",
			figure: shape.NewRectangle(10, 5),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// I don't quite understand yet how this works,
			// but looks cool anyway.
			require.Implements(t, (*shape.Shape)(nil), tC.figure)

			area := tC.figure.Area()
			str := tC.figure.String()

			switch tC.desc {
			case "Circle":
				require.EqualValues(t, 78.5375, area)
				require.Equal(t, "Circle: radius 5", str)
			case "Triangle":
				require.EqualValues(t, 24, area)
				require.Equal(t, "Triangle: base 8, height 6", str)
			case "Rectangle":
				require.EqualValues(t, 50, area)
				require.Equal(t, "Rectangle: length 10, width 5", str)
			default:
				require.Fail(t, "unknown case")
			}
		})
	}
}
