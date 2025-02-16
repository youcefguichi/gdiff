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

func lcs(s1, s2 []string) ([]string, map[int]int, map[int]int) {
	m := len(s1)
	n := len(s2)
	cur := make([]int, m+1)
	prev := make([]int, m+1)
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

func GenerateDiff(texta []string, textb []string, removed *map[int]int, inserted *map[int]int) []string {
	textaIdx := 0
	textbIdx := 0
	var diff []string

	if len(*removed) == 0 && len(*inserted) == 0 {
		return []string{"No changes"}
	}

	for {

		if textaIdx > len(texta)-1 && textbIdx > len(textb)-1 {
			break
		}

		if _, exists := (*removed)[textaIdx]; exists {
			diff = append(diff, fmt.Sprintf("%s- %s %s", red, string(texta[textaIdx]), reset))
		}

		if _, exists := (*inserted)[textbIdx]; exists {
			diff = append(diff, fmt.Sprintf("%s+ %s %s", green, string(textb[textbIdx]), reset))
		}

		textaIdx++
		textbIdx++
	}

	return diff
}

// func hashList(list []int) map[string]bool {

// }

func PrintDifff(diff, text1, text2 []string, removed map[int]int, inserted map[int]int, lineChangesTracker []int, depth int) {
	diffCurrentStartIndex := 0
	diffCurrentEndIndex := 0
	changeStarteIndex := -1
	changeEndIndex := 0

	for {

		if len(removed) == 0 && len(inserted) == 0 {
			break
		}

		// CurrentDiffStartIdx = lineChangesTracker[dIdx]
		// CurrentDiffEndIdx = lineChangesTracker[dIdx]


        // Calculate change width
		for i := changeStarteIndex + 1; i < len(diff); i++ {

			if _, exists := removed[i]; exists {
				changeEndIndex += 1
			}

			if _, exists := inserted[i-1]; exists{
				changeEndIndex += 1
			}
		}


		ctxStart := CurrentDiffStartIdx - depth
		ctxEnd := CurrentDiffEndIdx + depth

		if ctxStart < 0 {
			ctxStart = 0
		}

		if ctxEnd > max(len(text1), len(text2)) {
			ctxEnd = max(len(text1), len(text2))
		}
		diffIndex := dIdx
		for i := ctxStart; i < ctxEnd; i++ {

			if IndexExist(removed, i) {
				if diffIndex < len(diff) {
					fmt.Println(diff[diffIndex])
					diffIndex++
				}
			} else if IndexExist(inserted, i+1) {
				if diffIndex < len(diff) {
					fmt.Println(diff[diffIndex])
					diffIndex++
				}

			} else {
				if i < len(text1) {
					fmt.Println(text1[i])
				} else if i < len(text2) {
					fmt.Println(text2[i])
				} else {
					continue
				}

			}

		}

		for i, val := range lineChangesTracker {
			if val == CurrentDiffEndIdx {
				dIdx = i + 1
			}
		}
		if dIdx == len(lineChangesTracker) {
			break
		}

	}
}

func IndexExist(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
