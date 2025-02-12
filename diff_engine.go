package main

import (
	"bufio"
	"fmt"
	"os"
)

func createGrid(m, n int) [][]int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	return grid
}

func lcs(s1, s2 []string) ([]string, []int, []int) {
	m := len(s1) + 1
	n := len(s2) + 1
	grid := createGrid(m, n)
	var lcs []string
	var inserted []int
	var removed []int

	// Calculate lcs
	for i := range len(s1) {
		for j := range len(s2) {
			if s1[i] == s2[j] {
				grid[i+1][j+1] = grid[i][j] + 1
			}
			if s1[i] != s2[j] {
				grid[i+1][j+1] = max(grid[i][j+1], grid[i+1][j])
			}
		}
	}

	// Construct lcs
	i := m - 1
	j := n - 1
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = append(lcs, s1[i-1])
			i--
			j--
		} else if grid[i-1][j] > grid[i][j-1] {
			removed = append(removed, i-1)
			i--
		} else {
			inserted = append(inserted, j-1)
			j--
		}
	}

	for i > 0 {
		removed = append(removed, i-1)
		i--
	}

	for j > 0 {
		inserted = append(inserted, j-1)
		j--
	}

	// lcs = reverseSlice(lcs)
	return lcs, removed, inserted
}

func generateDiff(text1 []string, text2 []string) ([]string, []int, []int, []int) {
	text1Index := 0
	text2Index := 0
	var diff []string
	var lineChangesTracker []int
	_, removed, inserted := lcs(text1, text2)

	if len(removed) == 0 && len(inserted) == 0 {
		return []string{"No changes"}, []int{}, []int{}, []int{}
	}

	for {

		if text1Index > len(text1)-1 && text2Index > len(text2)-1 {
			break
		}

		if IndexExist(removed, text1Index) {
			diff = append(diff, fmt.Sprintf("\033[31m- %s\033[0m", string(text1[text1Index])))
			lineChangesTracker = append(lineChangesTracker, text1Index)
		}

		if IndexExist(inserted, text2Index) {
			diff = append(diff, fmt.Sprintf("\033[32m+ %s\033[0m", string(text2[text2Index])))
			lineChangesTracker = append(lineChangesTracker, text2Index)
		}

		text1Index++
		text2Index++
	}
	return diff, lineChangesTracker, removed, inserted
}

func PrintDiff(diff, text1, text2 []string, removed []int, inserted []int, lineChangesTracker []int, depth int) {
	dIdx := 0
	var CurrentDiffStartIdx int
	var CurrentDiffEndIdx int
	for {
		if len(lineChangesTracker) == 0 {
			break
		}
		CurrentDiffStartIdx = lineChangesTracker[dIdx]
		CurrentDiffEndIdx = lineChangesTracker[dIdx]

		for i := CurrentDiffStartIdx + 1; i < len(diff); i++ {
			if IndexExist(removed, i) {
				CurrentDiffEndIdx += 1
			}
			if IndexExist(inserted, i-1) {
				CurrentDiffEndIdx += 1
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
