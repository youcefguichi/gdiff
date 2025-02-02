package main

import (
	"reflect"
	"testing"
)

func TestLcs(t *testing.T) {

	tests := []struct {
		s1  string
		s2  string
		lcs string
	}{
		{"ABCDEF", "ABCDEF", "ABCDEF"},
		{"ABC", "XYZ", ""},
		{"AABCXY", "XYZ", "XY"},
		{"", "", ""},
		{"ABCD", "AC", "AC"},
	}

	for _, test := range tests {
		result := lcs(test.s1, test.s2)
		if !reflect.DeepEqual(result, test.lcs) {
			t.Errorf("lcs(%s, %s) expected %s, got %s", test.s1, test.s2, test.lcs, result)
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
