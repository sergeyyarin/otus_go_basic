package binarysearch

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
			collection: []int{0, 4, 5, 10, 12, 24, 37, 80, 800, 1000},
			value:      1,
		},
		"collection_holds_the_value": {
			collection: []int{0, 1, 3, 4, 5, 7, 9, 16, 18, 56},
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
