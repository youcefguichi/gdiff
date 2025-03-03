package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	GREEN = "\033[32m"
	RED   = "\033[31m"
	RESET = "\033[0m"
)

type Change struct {
	Idx  int
	Prev string
	Curr string
}

type DiffChecker struct {
	sourceText     []string
	revisedText    []string
	removed        map[int]int
	inserted       map[int]int
	changesTracker []int
	diff           []Change
	depth          int
}

func NewDiffChecker(sourceText, revisedText []string, depth int) *DiffChecker {
	return &DiffChecker{
		sourceText:     sourceText,
		revisedText:    revisedText,
		depth:          depth,
		removed:        make(map[int]int),
		inserted:       make(map[int]int),
		changesTracker: make([]int, 0),
		diff:           make([]Change, 0),
	}
}

func (d *DiffChecker) lcs(s1, s2 []string) {
	m := len(s1)
	n := len(s2)
	cur := make([]int, n+1)
	prev := make([]int, n+1)
	var lcs []string

	// Calculate lcs
	for i := 1; i < m+1; i++ {
		cur, prev = prev, cur
		for j := 1; j < n+1; j++ {

			if s1[i-1] == s2[j-1] {
				cur[j] = prev[j-1] + 1
			}
			if s1[i-1] != s2[j-1] {
				cur[j] = max(cur[j-1], prev[j])
			}
		}
	}

	// Construct lcs
	i := m
	j := n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = append(lcs, s1[i-1])
			i--
			j--
		} else if prev[j] == prev[j-1] {
			d.inserted[j-1] = 1
			j--
		} else {
			d.removed[i-1] = 1
			i--
		}
	}

	for i > 0 {
		d.removed[i-1] = 1
		i--
	}

	for j > 0 {
		d.inserted[j-1] = 1
		j--
	}

}

func (d *DiffChecker) GenerateDiff() {
	var sourceTextIdx int
	var revisedTextIdx int
	var trackerIndex int
	var change Change

	if len(d.removed) == 0 && len(d.inserted) == 0 {
		return
	}

	for {

		if sourceTextIdx > len(d.sourceText)-1 && revisedTextIdx > len(d.revisedText)-1 {
			break
		}

		if _, exists := d.removed[sourceTextIdx]; exists {
			change.Idx = sourceTextIdx
			change.Prev = fmt.Sprintf("%s- %s %s", RED, string(d.sourceText[sourceTextIdx]), RESET)
			d.changesTracker = append(d.changesTracker, sourceTextIdx)
			trackerIndex++
		}

		if _, exists := d.inserted[revisedTextIdx]; exists {
			change.Idx = revisedTextIdx
			change.Curr = fmt.Sprintf("%s+ %s %s", GREEN, string(d.revisedText[revisedTextIdx]), RESET)

			if trackerIndex > 0 && d.changesTracker[trackerIndex-1] != revisedTextIdx {
				d.changesTracker = append(d.changesTracker, revisedTextIdx)
				trackerIndex++
			}

			if len(d.changesTracker) == 0 && len(d.removed) == 0 {
				d.changesTracker = append(d.changesTracker, revisedTextIdx)
			}
		}

		if change.Prev != "" || change.Curr != "" {
			d.diff = append(d.diff, change)
		}

		sourceTextIdx++
		revisedTextIdx++
	}
}

func (d *DiffChecker) calculateConsecutiveChanges() (int, int, int) {

	changeStartIdx := d.changesTracker[0]
	changeEndIdx := d.changesTracker[0]
	count := 1

	if len(d.changesTracker) == 1 {
		return changeStartIdx, changeEndIdx, 0
	}

	for i, val := range d.changesTracker[:len(d.changesTracker)-1] {

		if val+1 == d.changesTracker[i+1] {
			changeEndIdx++
			count++

		} else {
			break
		}

	}

	return changeStartIdx, changeEndIdx, count
}

func (d *DiffChecker) calculateContextLines(changeStartIdx int, changeEndIdx int) (int, int) {

	contextChangeStartIdx := changeStartIdx - d.depth
	contextChangeEndIdx := changeEndIdx + d.depth

	if contextChangeStartIdx < 0 {
		contextChangeStartIdx = 0
	}

	if contextChangeEndIdx > len(d.revisedText) && changeEndIdx < len(d.revisedText) {
		contextChangeEndIdx = len(d.revisedText) - 1
	}

	if contextChangeEndIdx > len(d.revisedText) && changeEndIdx > len(d.revisedText) {
		contextChangeEndIdx = len(d.sourceText) - 1
	}

	return contextChangeStartIdx, contextChangeEndIdx

}

func overlap(a1, a2, b1, b2 int) bool {
	return a1 <= b2 && b1 <= a2
}

func mergeIndices(a1, a2, b1, b2 int) (int, int) {
	return min(a1, b1), max(a2, b2)
}

func (d *DiffChecker) printDiffWithContext(contextChangeStartIdx int, contextChangeEndIdx int, ctxLinesCache *[]int) {

	for j := contextChangeStartIdx; j <= contextChangeEndIdx; j++ {
		found := false
		for _, row := range d.diff {

			if row.Idx == j {

				if row.Curr != "" {
					fmt.Println(row.Curr)
				}

				if row.Prev != "" {
					fmt.Println(row.Prev)
				}

				found = true
				break
			}
		}

		if !found {

			if contextChangeEndIdx > len(d.revisedText) {
				fmt.Println(d.sourceText[j])

			} else {
				fmt.Println(d.revisedText[j])
			}
		}
	}
}

func (d *DiffChecker) start() {
	var changeStartIdx int
	var changeEndIdx int
	var nextChangeIdx int
	var ctxLinesCache []int
	var overlapStartIdx int
	var overlapEndIdx int
	var ctxLineStartIdx int
	var ctxLineEndIdx int

	d.lcs(d.sourceText, d.revisedText)
	d.GenerateDiff()

	if len(d.inserted) == 0 && len(d.removed) == 0 {
		return
	}

	for {

		if len(d.changesTracker) == 0 {
			break
		}

		for {

			changeStartIdx, changeEndIdx, nextChangeIdx = d.calculateConsecutiveChanges()
			ctxLineStartIdx, ctxLineEndIdx = d.calculateContextLines(changeStartIdx, changeEndIdx)
			ctxLinesCache = append(ctxLinesCache, ctxLineStartIdx)
			ctxLinesCache = append(ctxLinesCache, ctxLineEndIdx)

			if len(ctxLinesCache) > 2 && overlap(ctxLineStartIdx, ctxLineEndIdx, ctxLinesCache[0], ctxLinesCache[1]) {
				overlapStartIdx, overlapEndIdx = mergeIndices(ctxLineStartIdx, ctxLineEndIdx, ctxLinesCache[0], ctxLinesCache[1])
				ctxLinesCache = append(ctxLinesCache, overlapStartIdx)
				ctxLinesCache = append(ctxLinesCache, overlapEndIdx)
				ctxLinesCache = ctxLinesCache[len(ctxLinesCache)-2:]
			}

			if len(ctxLinesCache) > 2 && !overlap(ctxLineStartIdx, ctxLineEndIdx, ctxLinesCache[0], ctxLinesCache[1]) {
				ctxLinesCache = ctxLinesCache[:2]
				break
			}

			d.changesTracker = d.changesTracker[nextChangeIdx:]

			if len(d.changesTracker) == 1 && nextChangeIdx == 0 {
				d.changesTracker = d.changesTracker[:0]
				break
			}

			if len(d.changesTracker) == 0 {
				break
			}

		}

		for i := 0; i < len(ctxLinesCache); i += 2 {

			ctxLineStartIdx = ctxLinesCache[i]
			ctxLineEndIdx = ctxLinesCache[i+1]

			d.printDiffWithContext(ctxLineStartIdx, ctxLineEndIdx, &ctxLinesCache)
		}

		ctxLinesCache = ctxLinesCache[:0]

	}

}

func readFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
