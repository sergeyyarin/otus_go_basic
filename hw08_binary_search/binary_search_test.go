package binary_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	testCases := map[string]struct {
		collection []int
		value      int
	}{
		"empty_collection": {
			collection: []int{},
			value:      1,
		},
		"collection_does_not_hold_the_value": {
			collection: []int{37, 0, 4, 12, 800, 5, 10, 80, 24, 1000},
			value:      1,
		},
		"collection_holds_the_value": {
			collection: []int{16, 9, 7, 1, 56, 4, 18, 5, 0, 3},
			value:      1,
		},
	}
	for test, tc := range testCases {
		t.Run(test, func(t *testing.T) {
			found := BinarySearch(tc.collection, tc.value)
			switch test {
			case "empty_collection":
				require.False(t, found)
			case "collection_does_not_hold_the_value":
				require.False(t, found)
			case "collection_holds_the_value":
				require.True(t, found)
			default:
				t.Errorf("unknown case")
			}
		})
	}
}
