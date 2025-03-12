package main

import "fmt"

func main() {
	sourceFile := []string{"line 1", "line 2", "line 3", "line 4", "line 5"}
	revisedFile := []string{"line 2", "line 2", "line 3", "line 4", "line 5"}
	ndc := NewDiffChecker(sourceFile, revisedFile, 0)
	ndc.lcs(sourceFile, revisedFile)
	ndc.GenerateDiff()
	fmt.Println(ndc.diff)
	// ndc.start()
}
