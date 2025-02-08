package main

import (
	 "fmt"
	// "bufio"
	// "os"
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



func generateDiff(text1 []string, text2 []string) {
	_, removed, inserted := lcs(text1, text2)
        if len(removed) == 0 && len(inserted) == 0 {
			fmt.Println("No difference")
			
		} 
		// fmt.Printf("> ")
		for i := range text1 {
			if IndexExist(removed, i) {
				fmt.Printf("\033[31m-%s\033[0m \n", string(text1[i]))
			} else {
				fmt.Printf("%s \n", string(text1[i]))
			}
		}
		fmt.Printf("\n")
		// fmt.Printf("< ")
		for j := range text2 {
			if IndexExist(inserted, j) {
				fmt.Printf("\033[32m+%s\033[0m \n", string(text2[j]))
			} else {
				fmt.Printf("%s \n", string(text2[j]))
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

