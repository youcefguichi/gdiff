package main

import (
	//"fmt"
	"reflect"
	"testing"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		name        string
		sourceText  []string
		revisedText []string

		// expected
		lcs      []string
		inserted map[int]int
		removed  map[int]int
	}{
		{
			name:        "Identical strings",
			sourceText:  []string{"A", "B", "C", "D", "E", "F"},
			revisedText: []string{"A", "B", "C", "D", "E", "F"},
			lcs:         []string{"A", "B", "C", "D", "E", "F"},
			inserted:    map[int]int{},
			removed:     map[int]int{},
		},
		{
			name:        "Completely different strings",
			sourceText:  []string{"A", "B", "C"},
			revisedText: []string{"X", "Y", "Z"},
			lcs:         []string(nil),
			inserted:    map[int]int{0: 1, 1: 1, 2: 1},
			removed:     map[int]int{0: 1, 1: 1, 2: 1},
		},
		{
			name:        "Partial overlap",
			sourceText:  []string{"A", "A", "B", "C", "X", "Y"},
			revisedText: []string{"X", "Y", "Z"},
			lcs:         []string{"X", "Y"},
			inserted:    map[int]int{2: 1},
			removed:     map[int]int{0: 1, 1: 1, 2: 1, 3: 1},
		},
		{
			name:        "Empty strings",
			sourceText:  []string{},
			revisedText: []string{},
			lcs:         []string(nil),
			inserted:    map[int]int{},
			removed:     map[int]int{},
		},
		{
			name:        "Partial match with deletions",
			sourceText:  []string{"A", "B", "C", "D"},
			revisedText: []string{"A", "C"},
			lcs:         []string{"A", "C"},
			inserted:    map[int]int{},
			removed:     map[int]int{1: 1, 3: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ndc := NewDiffChecker(test.sourceText, test.revisedText, 0)
			ndc.lcs(test.sourceText, test.revisedText)

			if !reflect.DeepEqual(ndc.lcsOutput, test.lcs) {
				t.Errorf("lcs(%s, %s) expected  lcs: %v  got lcs: %v", test.sourceText, test.revisedText, test.lcs, ndc.lcsOutput)
			}

			if !reflect.DeepEqual(ndc.inserted, test.inserted) {
				t.Errorf("lcs(%s, %s) \n expected  inserted: %v \n got %v", test.sourceText, test.revisedText, test.inserted, ndc.inserted)
			}

			if !reflect.DeepEqual(ndc.removed, test.removed) {
				t.Errorf("lcs(%s, %s) expected deleted: %v, got %v", test.sourceText, test.revisedText, test.removed, ndc.removed)
			}
		})
	}
}

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

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element slice",
			input:    []string{"line 1"},
			expected: []string{"line 1"},
		},
		{
			name:     "Multiple element slice",
			input:    []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			expected: []string{"line 5", "line 4", "line 3", "line 2", "line 1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reversed := reverseSlice(test.input)
			if !reflect.DeepEqual(reversed, test.expected) {
				t.Errorf("reverseSlice(%v) expected %v, got %v", test.input, test.expected, reversed)
			}
		})
	}
}
func TestProcessNextChange(t *testing.T) {
	tests := []struct {
		name           string
		changesTracker []int
		depth          int
		sourceText     []string
		revisedText    []string
		expectedStart  int
		expectedEnd    int
	}{
		{
			name:           "Single change with no overlap",
			changesTracker: []int{2},
			depth:          1,
			sourceText:     []string{"line 1", "line 2", "line 3", "line 4"},
			revisedText:    []string{"line 1", "line 2", "line X", "line 4"},
			expectedStart:  1,
			expectedEnd:    3,
		},
		{
			name:           "Multiple consecutive changes",
			changesTracker: []int{2, 3, 4},
			depth:          1,
			sourceText:     []string{"line 1", "line 2", "line 3", "line 4", "line 5"},
			revisedText:    []string{"line 1", "line 2", "line X", "line Y", "line Z"},
			expectedStart:  1,
			expectedEnd:    4,
		},
		{
			name:           "Non-consecutive changes with overlap",
			changesTracker: []int{2, 5},
			depth:          2,
			sourceText:     []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6"},
			revisedText:    []string{"line 1", "line 2", "line X", "line 4", "line Y", "line Z"},
			expectedStart:  0,
			expectedEnd:    5,
		},
		{
			name:           "No changes",
			changesTracker: []int{},
			depth:          1,
			sourceText:     []string{"line 1", "line 2", "line 3"},
			revisedText:    []string{"line 1", "line 2", "line 3"},
			expectedStart:  0,
			expectedEnd:    0, // TODO: if no changes, should return None
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ndc := NewDiffChecker(test.sourceText, test.revisedText, test.depth)
			ndc.changesTracker = test.changesTracker
			startIdx, endIdx := ndc.proccessNextChange()

			if startIdx != test.expectedStart {
				t.Errorf("proccessNextChange() expected startIdx: %v, got: %v", test.expectedStart, startIdx)
			}

			if endIdx != test.expectedEnd {
				t.Errorf("proccessNextChange() expected endIdx: %v, got: %v", test.expectedEnd, endIdx)
			}
		})
	}
}
