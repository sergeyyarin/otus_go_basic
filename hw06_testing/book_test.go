package hw06testing

import (
	"testing"

	"github.com/sergeyyarin/otus_go_basic/hw06_testing/book"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	testCases := []struct {
		desc string
		cmpr book.Comparator
		rhs  book.Book
		lhs  book.Book
	}{
		{
			desc: "Empty",
			cmpr: *book.NewComparator(book.Year),
			lhs:  book.Book{},
			rhs:  book.Book{},
		},
		{
			desc: "NegativeComparisonValue",
			cmpr: *book.NewComparator(book.Year),
			lhs:  book.NewBook("t2", "a2", 700, 2, 2.0),
			rhs:  book.NewBook("t1", "a1", -700, 1, 1.0),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.cmpr.Compare(&tC.lhs, &tC.rhs)

			switch tC.desc {
			case "Empty":
				require.False(t, result)
			case "NegativeComparisonValue":
				require.True(t, result)
			default:
				require.Fail(t, "unknown case")
			}
		})
	}
}
