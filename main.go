package main

func main() {
	file1 := readFile("text1.txt")
	file2 := readFile("text2.txt")
	diff, lineChangesTracker, removed, inserted := generateDiff(file1, file2)
	PrintDiff(diff, file1, file2, removed, inserted, lineChangesTracker, 4)
}
