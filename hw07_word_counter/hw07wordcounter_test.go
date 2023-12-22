package wordcounter

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  map[string]int
	}{
		"empty input": {
			input: "",
			want:  map[string]int{},
		},
		"input with spaces": {
			input: "  foo bar  baz   foo",
			want: map[string]int{
				"foo": 2,
				"bar": 1,
				"baz": 1,
			},
		},
		"input with mixed spaces and punctuation": {
			input: "  foo;bar,baz..foo.  	     ,bar	;baz",
			want: map[string]int{
				"foo": 2,
				"bar": 2,
				"baz": 2,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := CountWords(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("not equal: got %v; want %v", got, tc.want)
			}
		})
	}
}
