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
	MINUS = "- "
	PLUS  = "+ "
)

type Cache struct {
	startIdx int
	endIdx   int
}

type Change struct {
	Idx  int
	Prev string
	Curr string
}

type DiffChecker struct {
	sourceText      []string
	revisedText     []string
	sourceTextSize  int
	revisedTextSize int
	lcsOutput       []string
	removed         map[int]int
	inserted        map[int]int
	changesTracker  []int
	diff            []Change
	depth           int
}

func NewDiffChecker(sourceText, revisedText []string, depth int) *DiffChecker {
	return &DiffChecker{
		sourceText:      sourceText,
		revisedText:     revisedText,
		sourceTextSize:  len(sourceText),
		revisedTextSize: len(revisedText),
		depth:           depth,
		removed:         make(map[int]int),
		inserted:        make(map[int]int),
		changesTracker:  make([]int, 0),
		diff:            make([]Change, 0),
	}
}

func (d *DiffChecker) lcs(s1, s2 []string) {
	m := len(s1)
	n := len(s2)
	cur := make([]int, n+1)
	prev := make([]int, n+1)

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
			d.lcsOutput = append(d.lcsOutput, s1[i-1])
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

	d.lcsOutput = reverseSlice(d.lcsOutput)
}

func (d *DiffChecker) GenerateDiff() {
	var sourceTextIdx int
	var revisedTextIdx int
	var trackerIndex int
	var change Change

	if !d.hasChanges() {
		return
	}

	for {

		change = Change{}

		if sourceTextIdx > len(d.sourceText)-1 && revisedTextIdx > len(d.revisedText)-1 {
			break
		}

		if _, exists := d.removed[sourceTextIdx]; exists {

			if _, exists := d.inserted[revisedTextIdx]; !exists {
				change.Curr = ""
			}

			change.Idx = sourceTextIdx
			change.Prev = RED + MINUS + string(d.sourceText[sourceTextIdx]) + RESET
			d.changesTracker = append(d.changesTracker, sourceTextIdx)
			trackerIndex++
		}

		if _, exists := d.inserted[revisedTextIdx]; exists {

			if _, exists := d.removed[revisedTextIdx]; !exists {
				change.Prev = ""
			}

			change.Idx = revisedTextIdx
			change.Curr = GREEN + PLUS + string(d.revisedText[revisedTextIdx]) + RESET

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

	if len(d.changesTracker) == 0 {
		return 0, 0, 0
	}

	if len(d.changesTracker) == 1 {
		return d.changesTracker[0], d.changesTracker[0], 0
	}

	changeStartIdx := d.changesTracker[0]
	changeEndIdx := d.changesTracker[0]
	nextChangeStartIdx := 1

	for i, val := range d.changesTracker[:len(d.changesTracker)-1] {

		if val+1 == d.changesTracker[i+1] {
			changeEndIdx++
			nextChangeStartIdx++

		} else {
			break
		}

	}

	return changeStartIdx, changeEndIdx, nextChangeStartIdx
}

func (d *DiffChecker) calculateContextLines(changeStartIdx int, changeEndIdx int) (int, int) {
	var ctxStart int
	var ctxEnd int

	ctxStart = changeStartIdx - d.depth
	ctxEnd = changeEndIdx + d.depth

	if ctxStart < 0 {
		ctxStart = 0
	}

	if ctxEnd > d.revisedTextSize-1 && changeEndIdx < d.revisedTextSize {
		ctxEnd = d.revisedTextSize - 1
	}

	if ctxEnd > d.revisedTextSize-1 && changeEndIdx > d.revisedTextSize {
		ctxEnd = len(d.sourceText) - 1
	}

	return ctxStart, ctxEnd

}

func (d *DiffChecker) printDiffWithContext(ctxStart int, ctxEnd int) {
	var found bool

	for j := ctxStart; j <= ctxEnd; j++ {
		found = false
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

			if ctxEnd > len(d.revisedText) {
				fmt.Println(d.sourceText[j])

			} else {
				fmt.Println(d.revisedText[j])
			}
		}

	}
}

func (d *DiffChecker) start() {
	d.lcs(d.sourceText, d.revisedText)
	d.GenerateDiff()

	if !d.hasChanges() {
		return
	}

	d.changeSummary()

	for len(d.changesTracker) > 0 {
		startIdx, endIdx := d.proccessNextChange()
		d.printDiffWithContext(startIdx, endIdx)
		separator()
	}

}

func (d *DiffChecker) proccessNextChange() (int, int) {
	var changeStartIdx int
	var changeEndIdx int
	var nextChangeIdx int
	var overlapStartIdx int
	var overlapEndIdx int
	var ctxStart int
	var ctxEnd int
	Cache := newCache()

	firstIteration := true
	for {

		changeStartIdx, changeEndIdx, nextChangeIdx = d.calculateConsecutiveChanges()
		ctxStart, ctxEnd = d.calculateContextLines(changeStartIdx, changeEndIdx)

		if !firstIteration && overlap(ctxStart, ctxEnd, Cache.startIdx, Cache.endIdx) {

			overlapStartIdx, overlapEndIdx = mergeIndices(ctxStart, ctxEnd, Cache.startIdx, Cache.startIdx)

			Cache.startIdx = overlapStartIdx
			Cache.endIdx = overlapEndIdx
		}

		if !firstIteration && !overlap(ctxStart, ctxEnd, Cache.startIdx, Cache.endIdx) {
			return Cache.startIdx, Cache.endIdx
		}

		if firstIteration {

			Cache.startIdx = ctxStart
			Cache.endIdx = ctxEnd
			firstIteration = false
		}

		d.changesTracker = d.changesTracker[nextChangeIdx:]

		if len(d.changesTracker) == 1 && nextChangeIdx == 0 {

			d.changesTracker = d.changesTracker[:0]
			return Cache.startIdx, Cache.endIdx
		}

		if len(d.changesTracker) == 0 {
			break
		}

	}

	return Cache.startIdx, Cache.endIdx
}

func (d *DiffChecker) hasChanges() bool {
	return !(len(d.inserted) == 0 && len(d.removed) == 0)
}

func loadFile(filename string) []string {
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

func newCache() *Cache {
	return &Cache{
		startIdx: 0,
		endIdx:   0,
	}
}

func overlap(a1, a2, b1, b2 int) bool {
	return a1 <= b2 && b1 <= a2
}

func mergeIndices(a1, a2, b1, b2 int) (int, int) {
	return min(a1, b1), max(a2, b2)
}

func reverseSlice(data []string) []string {
	for i := 0; i < len(data)/2; i++ {
		j := len(data) - 1 - i
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func separator() {
	fmt.Println()
	fmt.Println("\033[34m @@ Hidden Lines @@ \033[0m")
	fmt.Println()
}

func (d *DiffChecker) changeSummary() {
	fmt.Printf("\033[34m @@ change summary: %s+%d%s %s-%d%s @@ \033[0m\n", GREEN, len(d.inserted), RESET, RED, len(d.removed), RESET)
	fmt.Println()
}
