package main

import (
	//"fmt"
	"reflect"
	"testing"
)

// func TestLcs(t *testing.T) {

// 	tests := []struct {
// 		s1       string
// 		s2       string
// 		lcs      string
// 		removed  string
// 		inserted string
// 	}{
// 		{"ABCDEF", "ABCDEF", "ABCDEF", "", ""},
// 		{"ABC", "XYZ", "", "ABC", "XYZ"},
// 		{"AABCXY", "XYZ", "XY", "AABC", "Z"},
// 		{"", "", "", "", ""},
// 		{"ABCD", "AC", "AC", "BD", ""},
// 	}

// 	for _, test := range tests {
// 		lcs, removed, inserted := lcs(test.s1, test.s2)

// 		if !reflect.DeepEqual(lcs, test.lcs) {
// 			t.Errorf("lcs(%s, %s): expected '%s', got '%s'", test.s1, test.s2, test.lcs, lcs)
// 		}

// 		if !reflect.DeepEqual(removed, test.removed) {
// 			t.Errorf("lcs(%s, %s): expected to remove '%s' instead of '%s'", test.s1, test.s2, test.removed, removed)
// 		}

// 		if !reflect.DeepEqual(inserted, test.inserted) {
// 			t.Errorf("lcs(%s, %s): expected to insert '%s' instead of '%s'", test.s1, test.s2, test.inserted, inserted)
// 		}
// 	}

// }

// func TestReverseSlice(t *testing.T) {

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

// 	for _, test := range tests {
// 		result := reverseSlice(test.input)
// 		if !reflect.DeepEqual(result, test.expected) {
// 			t.Errorf("reverseSlice(%s) expected %s, got %s", test.input, test.expected, result)
// 		}
// 	}
// }

func TestGenerateDiff(t *testing.T) {
	tests := []struct {
		name        string
		sourceFile  []string
		revisedFile []string
		depth       int
		expected    []Change
	}{
		{
			name:        "No changes",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			depth:       0,
			expected:    []Change{},
		},
		{
			name:        "Modification",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 2", "line 2", "line 3", "line 4", "line 5"},
			depth:       0,
			expected:    []Change{{0, RED + MINUS + "line 1" + RESET, GREEN + PLUS + "line 2" + RESET}},
		},
		{
			name:        "Insertion",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6"},
			depth:       0,
			expected:    []Change{{5, "", GREEN + PLUS + "line 6" + RESET}},
		},
		{
			name:        "Deletion",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4"},
			depth:       0,
			expected:    []Change{{4, RED + MINUS + "line 5" + RESET, ""}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ndc := NewDiffChecker(test.sourceFile, test.revisedFile, test.depth)
			ndc.lcs(test.sourceFile, test.revisedFile)
			ndc.GenerateDiff()
			if !reflect.DeepEqual(ndc.diff, test.expected) {
				t.Errorf("generateDiff(%s, %s) expected %v, got %v", test.sourceFile, test.revisedFile, test.expected, ndc.diff)
			}
		})
	}
}
