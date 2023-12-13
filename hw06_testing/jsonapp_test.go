package hw06testing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/reader"
	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/types"
	"github.com/stretchr/testify/require"
)

func TestJsonAppReader(t *testing.T) {
	testCases := []struct {
		desc string
		file string
		data []types.Employee
	}{
		{
			desc: "ValidFilePath",
			file: "jsonapp/data.json",
			data: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
		{
			desc: "InvalidFilePath",
			file: "nowhere/data.json",
			data: []types.Employee{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := reader.ReadJSON(tC.file)
			require.True(t, cmp.Equal(tC.data, actual))
			switch tC.desc {
			case "ValidFilePath":
				require.NoError(t, err)
			case "InvalidFilePath":
				require.Contains(t, err.Error(), "no such file or directory")
			default:
				require.Fail(t, "unknown case")
			}
		})
	}
}
