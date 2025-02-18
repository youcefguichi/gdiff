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

func GenerateDiff(texta []string, textb []string, removed *map[int]int, inserted *map[int]int) ([]string, []int) {
	textaIdx := 0
	textbIdx := 0
	var diff []string
	var changesTracker []int
	i := -1

	if len(*removed) == 0 && len(*inserted) == 0 {
		return []string{"No changes"}, []int{}
	}

	for {

		if textaIdx > len(texta)-1 && textbIdx > len(textb)-1 {
			break
		}

		if _, exists := (*removed)[textaIdx]; exists {
			diff = append(diff, fmt.Sprintf("%s- %s %s", red, string(texta[textaIdx]), reset))
			changesTracker = append(changesTracker, textaIdx)
			i++
		}

		if _, exists := (*inserted)[textbIdx]; exists {
			diff = append(diff, fmt.Sprintf("%s+ %s %s", green, string(textb[textbIdx]), reset))
			if i > 0 && changesTracker[i-1] != textbIdx {
				changesTracker = append(changesTracker, textbIdx)
				i++
			}
		}

		textaIdx++
		textbIdx++
	}

	return diff, changesTracker
}

// func hashList(list []int) map[string]bool {

// }

func PrintDifff(diff, text1, text2 []string, removed map[int]int, inserted map[int]int, changesTracker []int, depth int) {
	var changeStartIdx int
	var changeEndIdx int
	lastChangeIteratedIndex := -1
    
	if len(inserted) == 0 && len(removed) == 0 {
		return
	}

	for changeEndIdx < len(diff) {
		changeStartIdx = lastChangeIteratedIndex + 1
		changeEndIdx = lastChangeIteratedIndex + 1

		// claculate consecutive changes width
		for idx, val := range changesTracker {

			// exist in both removed and inserted
			if _, existsR := removed[val]; existsR {
				if _, existsI := inserted[val]; existsI {
					changeEndIdx += 2
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

			if idx+1 < len(changesTracker) && changesTracker[idx] != changesTracker[idx+1] {
				lastChangeIteratedIndex = idx
				break
			}
		}

		// calculate context lines

		ctxLineStartIdx := changeStartIdx - depth
		ctxLineEndIdx := changeEndIdx + depth

		if ctxLineStartIdx < 0 {
			ctxLineStartIdx = 0
		}
		if ctxLineEndIdx > len(diff) {
			ctxLineEndIdx = len(text2)
		}

		for i := ctxLineStartIdx; i < ctxLineEndIdx; i++ {
			if changeStartIdx <= i && i < changeEndIdx {
				fmt.Println(diff[i])
			} else {
				fmt.Println(text2[i])
			}
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
