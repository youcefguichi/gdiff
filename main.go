package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: ./diff <source_file> <revised_file> <context_depth>")
		os.Exit(1)
	}

	sourceText := loadFile(os.Args[1])
	revisedText := loadFile(os.Args[2])
	contextDepth, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error: contextDepth must be an integer")
		os.Exit(1)
	}

	ndc := NewDiffChecker(sourceText, revisedText, contextDepth)
	ndc.start()
}
