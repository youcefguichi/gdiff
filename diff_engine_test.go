package main

import (
	"reflect"
	"testing"
)

func TestLcs(t *testing.T) {

	tests := []struct {
		s1       string
		s2       string
		lcs      string
		removed  string
		inserted string
	}{
		{"ABCDEF", "ABCDEF", "ABCDEF", "", ""},
		{"ABC", "XYZ", "", "ABC", "XYZ"},
		{"AABCXY", "XYZ", "XY", "AABC", "Z"},
		{"", "", "", "", ""},
		{"ABCD", "AC", "AC", "BD", ""},
	}

	for _, test := range tests {
		lcs, removed, inserted := lcs(test.s1, test.s2)

		if !reflect.DeepEqual(lcs, test.lcs) {
			t.Errorf("lcs(%s, %s): expected '%s', got '%s'", test.s1, test.s2, test.lcs, lcs)
		}

		if !reflect.DeepEqual(removed, test.removed) {
			t.Errorf("lcs(%s, %s): expected to remove '%s' instead of '%s'", test.s1, test.s2, test.removed, removed)
		}

		if !reflect.DeepEqual(inserted, test.inserted) {
			t.Errorf("lcs(%s, %s): expected to insert '%s' instead of '%s'", test.s1, test.s2, test.inserted, inserted)
		}
	}

}

func TestReverseSlice(t *testing.T) {

	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("ABCDEF"), []byte("FEDCBA")},
		{[]byte("ABC"), []byte("CBA")},
		{[]byte("AABCXY"), []byte("YXCBAA")},
		{[]byte("ABCD"), []byte("DCBA")},
		{[]byte(""), []byte("")},
	}

	for _, test := range tests {
		result := reverseSlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("reverseSlice(%s) expected %s, got %s", test.input, test.expected, result)
		}
	}
}

// func TestComputeAllLcs(t *testing.T) {
// 	tests := []struct {
// 		input    []byte
// 		expected []byte
// 	}{
// 		{[]byte("ABCDEF"), []byte("FEDCBA")},
// 		{[]byte("ABC"), []byte("CBA")},
// 		{[]byte("AABCXY"), []byte("YXCBAA")},
// 		{[]byte("ABCD"), []byte("DCBA")},
// 		{[]byte(""), []byte("")},
// 	}
// }
