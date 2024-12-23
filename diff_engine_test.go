package main

import (
	"reflect"
	"testing"
)

func TestLcs(t *testing.T) {

	tests := []struct {
		s1  string
		s2  string
		lcs int
	}{
		{"ABCDEF", "ABCDEF", 6},
		{"ABC", "XYZ", 0},
		{"AABCXY", "XYZ", 2},
		{"", "", 0},
		{"ABCD", "AC", 2},
	}

	for _, test := range tests {
		result := lcs(test.s1, test.s2)
		if !reflect.DeepEqual(result, test.lcs) {
			t.Errorf("lcs(%s, %s) expected %d, got %d", test.s1, test.s2, test.lcs, result)
		}
	}

}
