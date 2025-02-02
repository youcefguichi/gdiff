package main

import "fmt"

func main() {

	s1 := []string{"_, trackLineChanges := computeAllLcs(s1, s2)", "generateDiff(s1,s2,trackLineChanges)"}
	s2 := []string{"//_, trackLineChanges := computeAllLcs(s1, s2)", "//generateDiff(s1,s2,trackLineChanges)"}
	// s3 := []string{"Coding Challenges helps you become a better software engineer through that build real applications.",
	//      "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
	//      "I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
	//      "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities."}
	// s4 := []string{"Helping you become a better software engineer through coding challenges that build real applications.",
	//      "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
	//      "These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
	//      "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities."}

	// fmt.Println(lcs(s1, s2))
	// fmt.Println(s1, s2)
	// fmt.Printf("result %q", reverseSlice([]byte(s1)))
	computedLcs, trackLineChanges := computeAllLcs(s1, s2)
	generateDiff(s1, s2, trackLineChanges)
	// _, trackLineChanges = computeAllLcs(s3, s4)
	// generateDiff(s3,s4,trackLineChanges)
	fmt.Printf("result %q \n", computedLcs)
	fmt.Printf("result %d \n", trackLineChanges)

	// Lines 1: "This is a test which contains:", "this is the lcs"
	// Lines 2: "this is the lcs", "we're testing"
	// Expected LCS: "this is the lcs"
}

// s1 := []string{"Coding Challenges helps you become a better software engineer through that build real applications.",
//          "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
//          "I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
//          "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities."}
// 	s2 := []string{"Helping you become a better software engineer through coding challenges that build real applications.",
//          "I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
//          "These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
//          "Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities."}
