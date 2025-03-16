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
			name:        "Single Modification",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 2", "line 2", "line 3", "line 4", "line 5"},
			depth:       0,
			expected:    []Change{{0, RED + MINUS + "line 1" + RESET, GREEN + PLUS + "line 2" + RESET}},
		},
		{
			name:        "Multiple Modifications",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 34", "line 345", "line 4", "line 5"},
			depth:       0,
			expected: []Change{
				{1, RED + MINUS + "line 2" + RESET, GREEN + PLUS + "line 34" + RESET},
				{2, RED + MINUS + "line 3" + RESET, GREEN + PLUS + "line 345" + RESET},
			},
		},
		{
			name:        "Insertion",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6"},
			depth:       0,
			expected:    []Change{{5, "", GREEN + PLUS + "line 6" + RESET}},
		},
		{
			name:        "Multiple Insertions",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6", "line 7", "line 8"},
			depth:       0,
			expected: []Change{{5, "", GREEN + PLUS + "line 6" + RESET},
				{6, "", GREEN + PLUS + "line 7" + RESET},
				{7, "", GREEN + PLUS + "line 8" + RESET},
			},
		},
		{
			name:        "Deletion",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2", "line 3", "line 4"},
			depth:       0,
			expected:    []Change{{4, RED + MINUS + "line 5" + RESET, ""}},
		},

		{
			name:        "Multiple Deletions",
			sourceFile:  []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedFile: []string{"line 1", "line 2"},
			depth:       0,
			expected: []Change{
				{2, RED + MINUS + "line 3" + RESET, ""},
				{3, RED + MINUS + "line 4" + RESET, ""},
				{4, RED + MINUS + "line 5" + RESET, ""},
			},
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

func TestCalcualteConsecutiveChanges(t *testing.T) {
	tests := []struct {
		name           string
		changesTracker []int

		// expected
		changeStartIdx     int
		changeEndIdx       int
		nextChangeStartIdx int
	}{
		{
			name:               "No changes",
			changesTracker:     []int{},
			changeStartIdx:     0,
			changeEndIdx:       0,
			nextChangeStartIdx: 0,
		},
		{
			name:               "Consec Changes",
			changesTracker:     []int{0, 1, 2, 3, 4, 5, 6},
			changeStartIdx:     0,
			changeEndIdx:       6,
			nextChangeStartIdx: 7,
		},
		{
			name:               "Non Consec Changes",
			changesTracker:     []int{0, 1, 2, 10, 24, 55, 45},
			changeStartIdx:     0,
			changeEndIdx:       2,
			nextChangeStartIdx: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sourceText := []string{}
			revisedText := []string{}
			ndc := NewDiffChecker(sourceText, revisedText, 0)
			ndc.changesTracker = test.changesTracker
			changeStartIdx, changeEndIdx, nextChangeStartIdx := ndc.calculateConsecutiveChanges()

			if !reflect.DeepEqual(changeStartIdx, test.changeStartIdx) {
				t.Errorf("calculateConsecutiveChanges() expected changeStartIdx: %v, got: %v", test.changeStartIdx, changeStartIdx)
			}

			if !reflect.DeepEqual(changeEndIdx, test.changeEndIdx) {
				t.Errorf("calculateConsecutiveChanges() expected changeEndIdx: %v, got: %v", test.changeEndIdx, changeEndIdx)
			}

			if !reflect.DeepEqual(nextChangeStartIdx, test.nextChangeStartIdx) {
				t.Errorf("calculateConsecutiveChanges() expected nextChangeIdx: %v, got: %v", test.nextChangeStartIdx, nextChangeStartIdx)
			}
		})
	}
}

func TestCalculateContextLines(t *testing.T) {
	tests := []struct {
		name            string
		changeStartIdx  int
		changeEndIdx    int
		depth           int
		revisedTextSize int
		sourceTextSize  int

		// expected output
		ctxStart int
		ctxEnd   int
	}{
		{
			name:            "ctx end > revisedTextSize && ctx start < 0",
			changeStartIdx:  0,
			changeEndIdx:    0,
			depth:           7,
			revisedTextSize: 5,
			sourceTextSize:  5,

			// expected output
			ctxStart: 0,
			ctxEnd:   4,
		},
		{
			name:            "Change within the text size range",
			changeStartIdx:  2,
			changeEndIdx:    4,
			depth:           2,
			revisedTextSize: 10,
			sourceTextSize:  10,

			// expected output
			ctxStart: 0,
			ctxEnd:   6,
		},
		{
			name:            "Behaviour on reduced revised text size",
			changeStartIdx:  2,
			changeEndIdx:    4,
			depth:           10,
			revisedTextSize: 7,
			sourceTextSize:  10,

			// expected output
			ctxStart: 0,
			ctxEnd:   6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ndc := NewDiffChecker([]string{}, []string{}, test.depth)
			ndc.revisedTextSize = test.revisedTextSize
			ndc.sourceTextSize = test.sourceTextSize
			ctxStart, ctxEnd := ndc.calculateContextLines(test.changeStartIdx, test.changeEndIdx)

			if !reflect.DeepEqual(ctxStart, test.ctxStart) {
				t.Errorf("calculateContextLines() expected ctxStart: %v, got: %v", test.ctxStart, ctxStart)
			}

			if !reflect.DeepEqual(ctxEnd, test.ctxEnd) {
				t.Errorf("calculateContextLines() expected ctxEnd: %v, got: %v", test.ctxEnd, ctxEnd)
			}
		})
	}
}
