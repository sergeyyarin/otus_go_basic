package hw06testing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/printer"
	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/reader"
	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/types"
	"github.com/stretchr/testify/require"
)

func TestJsonAppReader(t *testing.T) {
	testCases := []struct {
		desc     string
		file     string
		expected []types.Employee
	}{
		{
			desc: "ValidFilePath",
			file: "jsonapp/data.json",
			expected: []types.Employee{
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
			desc:     "InvalidFilePath",
			file:     "nowhere/data.json",
			expected: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := reader.ReadJSON(tC.file)
			require.True(t, cmp.Equal(tC.expected, actual))
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

func TestJsonAppPrinter(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []types.Employee
		expected string
	}{
		{
			desc:     "EmptyInput",
			input:    []types.Employee{},
			expected: "",
		},
		{
			desc: "ValidInput",
			input: []types.Employee{
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
			expected: "User ID: 10; Age: 25; Name: Rob; Department ID: 3; User ID: 11; Age: 30; Name: George; Department ID: 2; ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := printer.PrintStaff(tC.input)
			require.Equal(t, tC.expected, actual)
		})
	}
}
