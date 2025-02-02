package main

import "fmt"

func createGrid(m, n int) [][]int {

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	return grid
}

func lcs(s1, s2 string) string {
	m := len(s1) + 1
	n := len(s2) + 1
	grid := createGrid(m, n)
	var lcs []byte

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
			i--
		} else {
			j--
		}
	}
	lcs = reverseSlice(lcs)
	return string(lcs)
}

func reverseSlice(data []byte) []byte {
	j := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
	return data
}

func computeAllLcs(s1, s2 []string) ([]string, []int) {
	var computedLcs []string
	var trackLineChanges []int
	for i := range s1 {
		lcs := lcs(s1[i], s2[i])
		if lcs != s2[i] {
			trackLineChanges = append(trackLineChanges, 1)
		} else {
			trackLineChanges = append(trackLineChanges, 0)
		}
		computedLcs = append(computedLcs, lcs)
	}
	return computedLcs, trackLineChanges
}

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func generateDiff(s1, s2 []string, lineTracker []int) {
	for i := range s1 {
		if lineTracker[i] == 1 {
			fmt.Printf("%s> %q%s\n", Red, s1[i], Reset)
			fmt.Printf("%s< %q%s\n", Green, s2[i], Reset)
		}
	}
}

