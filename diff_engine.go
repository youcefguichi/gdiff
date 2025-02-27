package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

type DiffItem struct {
	Idx   int
	Lines []string
}

func lcs(s1, s2 []string) ([]string, map[int]int, map[int]int) {
	m := len(s1)
	n := len(s2)
	cur := make([]int, n+1)
	prev := make([]int, n+1)
	var lcs []string
	inserted := make(map[int]int)
	removed := make(map[int]int)

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
			inserted[j-1] = 1
			j--
		} else {
			removed[i-1] = 1
			i--
		}
	}

	for i > 0 {
		removed[i-1] = 1
		i--
	}

	for j > 0 {
		inserted[j-1] = 1
		j--
	}

	return lcs, removed, inserted
}

func GenerateDiff(texta []string, textb []string, removed *map[int]int, inserted *map[int]int) ([]DiffItem, []int) {
	textaIdx := 0
	textbIdx := 0
	var diff []DiffItem
	var changesTracker []int
	// i := 0

	if len(*removed) == 0 && len(*inserted) == 0 {
		return []DiffItem{}, []int{}
	}
	trackerIndex := 0
	for {

		if textaIdx > len(texta)-1 && textbIdx > len(textb)-1 {
			break
		}
		diffItem := DiffItem{}
		if _, exists := (*removed)[textaIdx]; exists {
			diffItem.Idx = textaIdx
			diffItem.Lines = append(diffItem.Lines, fmt.Sprintf("%s- %s %s", red, string(texta[textaIdx]), reset))
			changesTracker = append(changesTracker, textaIdx)
			trackerIndex++
		}

		if _, exists := (*inserted)[textbIdx]; exists {
			diffItem.Idx = textbIdx
			diffItem.Lines = append(diffItem.Lines, fmt.Sprintf("%s+ %s %s", green, string(textb[textbIdx]), reset))

			if trackerIndex > 0 && changesTracker[trackerIndex-1] != textbIdx {
				changesTracker = append(changesTracker, textbIdx)
				trackerIndex++
			}

			if len(changesTracker) == 0 && len(*removed) == 0 {
				changesTracker = append(changesTracker, textbIdx)
			}
		}

		if len(diffItem.Lines) != 0 {
			diff = append(diff, diffItem)
		}
		textaIdx++
		textbIdx++
	}

	return diff, changesTracker
}

func calculateConsecutiveChanges(changesTracker []int) (int, int, int) {

	changeStartIdx := changesTracker[0]
	changeEndIdx := changesTracker[0]
	count := 1

	if len(changesTracker) == 1 {
		return changeStartIdx, changeEndIdx, 0
	}

	for i, val := range changesTracker[:len(changesTracker)-1] {

		if val+1 == changesTracker[i+1] {
			changeEndIdx++
			count++

		} else {
			break
		}

	}

	return changeStartIdx, changeEndIdx, count
}

func calculateContextLines(changeStartIdx int, changeEndIdx int, text1, text2 []string, depth int) (int, int) {

	ctxLineStartIdx := changeStartIdx - depth
	ctxLineEndIdx := changeEndIdx + depth

	if ctxLineStartIdx < 0 {
		ctxLineStartIdx = 0
	}

	if ctxLineEndIdx > len(text2) && changeEndIdx < len(text2) {
		ctxLineEndIdx = len(text2) - 1
	}

	if ctxLineEndIdx > len(text2) && changeEndIdx > len(text2) {
		ctxLineEndIdx = len(text1) - 1
	}

	return ctxLineStartIdx, ctxLineEndIdx

}

func overlap(a1, a2, b1, b2 int) bool {
	return a1 <= b2 && b1 <= a2
}

func mergeIndices(a1, a2, b1, b2 int) (int, int) {
	return min(a1, b1), max(a2, b2)
}

func displayDiffWithCtxLines(ctxLineStartIdx int, ctxLineEndIdx int, diff []DiffItem, text1, text2 []string, ctxLinesCache *[]int) {

	for j := ctxLineStartIdx; j <= ctxLineEndIdx; j++ {
		found := false
		for _, row := range diff {

			if row.Idx == j {
				for _, line := range row.Lines {
					fmt.Println(line)
				}
				found = true
				break
			}

		}

		if !found {
			if ctxLineEndIdx > len(text2) && ctxLineEndIdx < len(text1) {
				fmt.Println(text1[j])
			}

			if ctxLineEndIdx < len(text2) && ctxLineEndIdx < len(text1) {
				fmt.Println(text2[j])
			}

		}
	}

}

func PrintDifff(diff []DiffItem, text1, text2 []string, removed map[int]int, inserted map[int]int, changesTracker []int, depth int) {
	var changeStartIdx int
	var changeEndIdx int
	var nextChangeIdx int
	var ctxLinesCache []int
	var overlapStartIdx int
	var overlapEndIdx int

	if len(inserted) == 0 && len(removed) == 0 {
		return
	}

	for {

		if len(changesTracker) == 0 {
			break
		}

		changeStartIdx, changeEndIdx, nextChangeIdx = calculateConsecutiveChanges(changesTracker)

		ctxLineStartIdx, ctxLineEndIdx := calculateContextLines(
			changeStartIdx,
			changeEndIdx,
			text1,
			text2,
			depth,
		)
		ctxLinesCache = append(ctxLinesCache, ctxLineStartIdx)
		ctxLinesCache = append(ctxLinesCache, ctxLineEndIdx)

		if overlap(ctxLineStartIdx, ctxLineEndIdx, ctxLinesCache[0], ctxLinesCache[1]) {

			if len(ctxLinesCache) > 2 {
				overlapStartIdx, overlapEndIdx = mergeIndices(ctxLineStartIdx, ctxLineEndIdx, ctxLinesCache[0], ctxLinesCache[1])
				ctxLinesCache = append(ctxLinesCache, overlapStartIdx)
				ctxLinesCache = append(ctxLinesCache, overlapEndIdx)
				ctxLinesCache = ctxLinesCache[:len(ctxLinesCache)-2]
				changesTracker = changesTracker[1:]
			}

			changesTracker = changesTracker[nextChangeIdx:]
			ctxLinesCache = ctxLinesCache[len(ctxLinesCache)-2:]
			
			if len(changesTracker) != 0 {
				continue
			}
		}

		displayDiffWithCtxLines(ctxLineStartIdx,
			ctxLineEndIdx,
			diff,
			text1,
			text2,
			&ctxLinesCache,
		)
        
		// changesTracker = changesTracker[nextChangeIdx:]
		// ctxLinesCache = ctxLinesCache[len(ctxLinesCache)-2:]

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
