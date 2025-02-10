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
		CurrentDiffStartIdx = lineChangesTracker[dIdx]
		CurrentDiffEndIdx = lineChangesTracker[dIdx]
		for i := CurrentDiffStartIdx + 1; i < len(diff); i++ {
			if IndexExist(removed, i) {
				CurrentDiffEndIdx += 1
			}
			if IndexExist(inserted, i) {
				CurrentDiffEndIdx += 1
			}
		}
		dIdx = CurrentDiffEndIdx + 1

		for i := CurrentDiffStartIdx - depth; i <= CurrentDiffEndIdx+depth; i++ {
			if IndexExist(removed, i) || IndexExist(inserted, i) {
				fmt.Println(diff[i])
			}

		}

	}
}

// func PrintContextLines(diff []string, currentLine string, currentLineIndex int,lineChangesTracker []int, text1 []string, text2 []string , removed []int, inserted []int, depth int) {

// 	startingIndex := currentLineIndex - depth
// 	endingIndex := currentLineIndex + depth

// 	if startingIndex < 0 {
// 		startingIndex = 0
// 	}

// 	if endingIndex > len(text1) {
// 		endingIndex = len(text1)
// 	}

// 	for dIdx,diffLine := range diff{

//     // calculate context for change
//     CalculateContextIds(dIdx, lineChangesTracker)

// 	for i := startingIndex; i < endingIndex ; i++{

//         if IndexExist(removed, i) {
//             fmt.Println(text1[i])
// 		}
// 		if IndexExist(inserted, i) {
// 			fmt.Println(text2[i])
// 		}
// 		if i == lineChangesTracker[dIdx]{
// 			fmt.Println(diffLine)
// 		}

// 		// if i == currentLineIndex && currentLine[0] == '-'{
// 			// fmt.Println(currentLine)
// 		// }
// 	}

// 	}

// }

// func CalculateContextIds(dIdx int, currentLineIdx int, startingIndex int, endingIndex int, lineChangesTracker []int, depth int) bool{

// }
// if currentLine[0] == '+' {

// 	}

// func upcomingChnages(lineChangesTracker []int) bool{
// 	for i := range(lineChangesTracker){
// 		if i == lineChangesTracker[i]{
// 			return false
// 		}
// 		return true
// 	}
// }

func PrintNextLines(text []string, currentIndex int, linesNum int) {
	for i := currentIndex; i < currentIndex+linesNum; i++ {
		if i > len(text)-1 {
			break
		}
		fmt.Println(text[i])
	}
}

func PrintPrevLines(text []string, currentIndex int, linesNum int) {
	for i := currentIndex; i > currentIndex-linesNum; i-- {
		if i < 0 {
			break
		}
		fmt.Println(text[i])
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
