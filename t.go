package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(500 * time.Millisecond)
	}
}

func PrintDiff(diff, text1, text2 []string, removed map[int]int, inserted map[int]int, lineChangesTracker []int, depth int) {
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