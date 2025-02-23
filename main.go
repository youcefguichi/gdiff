package main

// "fmt"
//"time"

func main() {
	file1 := readFile("text1.txt")
	file2 := readFile("text2.txt")
	_, removed, inserted := lcs(file1, file2)
	diff, changesTracker := GenerateDiff(file1, file2, &removed, &inserted)
	// fmt.Println(diff)
	PrintDifff(diff, file1, file2, removed, inserted, changesTracker, 7)
	// for _, line := range diff {
	// 	fmt.Println(line)
	// }
}
