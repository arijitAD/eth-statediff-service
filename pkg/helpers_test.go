package statediff_test

import (
	sd "github.com/vulcanize/eth-statediff-service/pkg"
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	// precond: a, b are sorted w/ no duplicates
	testCases := []struct {
		a, b, expected []string
	}{
		{[]string{}, []string{}, []string{}},
		{[]string{""}, []string{""}, []string{""}},
		{
			[]string{"a", "b"},
			[]string{"a", "c"},
			[]string{"a"},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
	}

	for _, test := range testCases {
		intersection := sd.FindIntersection(test.a, test.b)
		if !reflect.DeepEqual(test.expected, intersection) {
			t.Errorf("failed: TestFindIntersection\nexpected: %v\nactual %v\n", test.expected, intersection)
		}
	}
}
