package main

import (
	"flag"
	"fmt"
)

func main() {
	sourceText := flag.String("source-text", "", "Path to the source text file")
	revisedText := flag.String("revised-text", "", "Path to the revised text file")
	contextSize := flag.Int("context-size", 3, "Number of context lines")

	flag.Parse()

	if *sourceText == "" || *revisedText == "" {
		fmt.Println("Both --source-text and --revised-text flags are required")
		return
	}

	sourceFile := loadFile(*sourceText)
	revisedFile := loadFile(*revisedText)

	ndc := NewDiffChecker(sourceFile, revisedFile, *contextSize)
	ndc.start()
}
