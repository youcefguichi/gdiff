package main

import (
	"fmt"
	//"time"
)

func main() {
	file1 := readFile("text1.txt")
	file2 := readFile("text2.txt")
	_, removed, inserted := lcs(file1, file2)
	diff := generateDiff(file1, file2, &removed, &inserted)
	for _, line := range diff {
		fmt.Println(line)
	}
}
