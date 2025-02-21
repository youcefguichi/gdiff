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
			//changesTracker = append(changesTracker, textbIdx)
			// tempEntry := map[int]string{textbIdx: fmt.Sprintf("%s+ %s %s", green, string(textb[textbIdx]), reset)}
			// diff = append(diff, tempEntry)
			if trackerIndex > 0 && changesTracker[trackerIndex-1] != textbIdx {
				changesTracker = append(changesTracker, textbIdx)
				trackerIndex++
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

// func hashList(list []int) map[string]bool {

// }
func calculateConsecutiveChanges(changesTracker []int, changeStartIdx int, changeEndIdx int, removed map[int]int, inserted map[int]int) (int, int) {

	for _, val := range changesTracker {

		// exist in both removed and inserted
		if _, existsR := removed[val]; existsR {
			if _, existsI := inserted[val]; existsI {
				changeEndIdx += 1
			}
		}

		// exist in removed only and not in inserted
		if _, existsR := removed[val]; existsR {
			if _, existsI := inserted[val]; !existsI {
				changeEndIdx += 1
			}
		}

		// !exist in removed only and exist in inserted
		if _, existsR := removed[val]; !existsR {
			if _, existsI := inserted[val]; existsI {
				changeEndIdx += 1
			}
		}

		// if idx+1 < len(changesTracker) && changesTracker[idx] != changesTracker[idx+1] {
		// 	lastChangeIteratedIndex = idx
		// 	break
		// }
	}
	return changeStartIdx, changeEndIdx
}

func calculateContextLines(changeStartIdx int, changeEndIdx int, text2 []string, depth int) (int, int) {

	ctxLineStartIdx := changeStartIdx - depth
	ctxLineEndIdx := changeEndIdx + depth

	if ctxLineStartIdx < 0 {
		ctxLineStartIdx = 0
	}

	if ctxLineEndIdx > len(text2) {
		ctxLineEndIdx = len(text2)
	}

	return ctxLineStartIdx, ctxLineEndIdx

}

func displayDiffWithCtxLines(ctxLineStartIdx int, ctxLineEndIdx int, diff []DiffItem, text2 []string) {

	for j := ctxLineStartIdx; j < ctxLineEndIdx; j++ {
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
			fmt.Println(text2[j])
		}
	}

}
func findNextChangeSequence(diff []DiffItem, Idx int, changesTracker []int) (int, int, []int) {
	for i, row := range diff {
		if i+1 >= len(diff) {
			break
		}

		if row.Idx+1 == diff[i+1].Idx {
			Idx++
		} else {
			break
		}

	}
	// if Idx >= len(diff) {
	// 	break
	// }
	changeStartIdx := diff[Idx].Idx
	changeEndIdx := diff[Idx].Idx
	changesT := changesTracker[Idx:]

	return changeStartIdx, changeEndIdx, changesT
}

func PrintDifff(diff []DiffItem, text1, text2 []string, removed map[int]int, inserted map[int]int, changesTracker []int, depth int) {
	var changeStartIdx int
	var changeEndIdx int

	if len(inserted) == 0 && len(removed) == 0 {
		return
	}

	changeStartIdx = diff[0].Idx
	changeEndIdx = diff[0].Idx
	Idx := 0
	for changeStartIdx <= diff[len(diff)-1].Idx {

		changeStartIdx, changeEndIdx = calculateConsecutiveChanges(
			changesTracker,
			changeStartIdx,
			changeEndIdx,
			removed,
			inserted,
		)

		ctxLineStartIdx, ctxLineEndIdx := calculateContextLines(
			changeStartIdx,
			changeEndIdx,
			text2,
			depth,
		)

		displayDiffWithCtxLines(ctxLineStartIdx, ctxLineEndIdx, diff, text2) // i := ctxLineStartIdx
		Idx++
        
		findNextChangeSequence(diff, Idx, changesTracker)
		

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
