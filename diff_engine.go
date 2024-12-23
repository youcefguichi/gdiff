package main

// import "fmt"

func createGrid(m, n int) [][]int {

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	return grid
}

func lcs(s1, s2 string) (int, string) {
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
			i-- // Move up
		} else {
			j-- // Move left
		}
	}
	lcs = reverse(lcs)
	return grid[m-1][n-1], string(lcs)
}

func reverse(data []byte) []byte {
	for i := 0; i < len(data); i++ {
		j := len(data) - 1
		data[i], data[j] = data[j], data[i]
	}
	return data
}
