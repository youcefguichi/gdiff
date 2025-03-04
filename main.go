package main

func main() {
	file1 := readFile("text1.txt")
	file2 := readFile("text2.txt")
	ndc := NewDiffChecker(file1, file2, 3)
	ndc.start()
}
