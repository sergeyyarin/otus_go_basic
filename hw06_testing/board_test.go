package hw06testing

import (
	"testing"

	"github.com/sergeyyarin/otus_go_basic/hw06_testing/board"

	"github.com/stretchr/testify/require"
)

func TestBuildBoard(t *testing.T) {
	testCases := []struct {
		desc   string
		size   int
		expect string
	}{
		{
			desc:   "Negative",
			size:   -1,
			expect: "",
		},
		{
			desc:   "Empty",
			size:   0,
			expect: "",
		},
		{
			desc:   "Odd",
			size:   1,
			expect: "_\n",
		},
		{
			desc:   "Even",
			size:   2,
			expect: "_#\n#_\n",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			str, err := board.Build(tC.size)

			switch tC.desc {
			case "Negative", "Empty", "Odd":
				require.ErrorIs(t, err, board.ErrInvalidValue)
			case "Even":
				require.NoError(t, err)
				require.Equal(t, tC.expect, str)
			default:
				require.Fail(t, "unknown case")
			}
		})
	}
}
