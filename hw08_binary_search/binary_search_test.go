package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	testCases := map[string]struct {
		collection []int
		value      int
		expected   int
	}{
		"empty_collection": {
			collection: []int{},
			value:      1,
			expected:   -1,
		},
		"collection_does_not_hold_the_value": {
			collection: []int{0, 4, 5, 10, 12, 24, 37, 80, 800, 1000},
			value:      1,
			expected:   -1,
		},
		"collection_holds_the_value": {
			collection: []int{0, 1, 3, 4, 5, 7, 9, 16, 18, 56},
			value:      18,
			expected:   8,
		},
	}
	for test, tc := range testCases {
		t.Run(test, func(t *testing.T) {
			index := BinarySearch(tc.collection, tc.value)
			switch test {
			case "empty_collection":
				require.Equal(t, tc.expected, index)
			case "collection_does_not_hold_the_value":
				require.Equal(t, tc.expected, index)
			case "collection_holds_the_value":
				require.Equal(t, tc.expected, index)
			default:
				t.Errorf("unknown case")
			}
		})
	}
}
