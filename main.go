package main

import (
	"fmt"
	"time"
)

func main() {
	file1 := readFile("text1.txt")
	file2 := readFile("text2.txt")
	startTime := time.Now().UnixMilli()
	diff, lineChangesTracker, removed, inserted := generateDiff(file1, file2)
	PrintDiff(diff, file1, file2, removed, inserted, lineChangesTracker, 4)
	endTime := time.Now().UnixMilli()
	fmt.Printf("Execution time: %d ms\n", endTime-startTime)
}
