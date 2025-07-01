package main

import "fmt"

func PrintSlice(s []int) {
	fmt.Println(s)
}

func main() {
	s1 := []int{10, 20, 10}
	s2 := make([]int, 0, 10)

	// Didn't append
	Append(s1)
	Append(s2)

	PrintSlice(s1)
	PrintSlice(s2)

	// Append
	AppendPointer(&s1)
	AppendPointer(&s2)

	PrintSlice(s1)
	PrintSlice(s2)

	// printPointerSlice(&s1)
}

// to-do Print address
func Append(s []int) {
	s = append(s, 40)
}

// Passing pointer
func AppendPointer(s *[]int) {
	*s = append(*s, 30)
}

// What is happening here ??
func printPointerSlice(s *[]int) {

	// a is a copy of orignal s[i]
	for _, a := range *s {
		fmt.Printf("%d ", a)
	}
	fmt.Println()

	// original value of s[i]
	// dereferencing using (*s)
	for i := range *s {
		fmt.Printf("%d ", (*s)[i])
	}
	fmt.Println()

	// append another slice
	
	// pointer slice
	// var s []*int
}
