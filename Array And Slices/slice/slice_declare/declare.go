package main

import "fmt"

func PrintSlice(s []int) {
	fmt.Println(s)
}

func main() {
	var s1 []int  // nil slice
	s2 := []int{} // empty slice

	PrintSlice(s1)
	PrintSlice(s2)

	// make(type, length, capacity)
	s3 := make([]int, 0)
	s4 := make([]int, 0, 10)

	PrintSlice(s3)
	PrintSlice(s4)

	// s[i : j] -> i to j - 1
	// 0 <= i <= j <= cap(s)
	// slice beyond cap(s) causes panic
	// slice beyond len(s) extends the slice
	// length = j - i
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	s5 := arr[2:4]
	PrintSlice(s5)

	// cap(s) = 8
	// len(s) = 5
	s := []int{0, 1, 2, 3}
	s = append(s, 4)

	// Within cap, beyond len
	s6 := s[2:6]
	PrintSlice(s6)

	// Beyond cap, panic
	// s7 := s[2 : 10]
	// PrintSlice(s7) // Panic

	// len(s) = 2
	// cap(s) = 4
	// len(s8) = 1
	// cap(s8) = 3 = cap(s) - i
	// s8 = s[i : j]
	s = make([]int, 2, 4)
	s8 := s[1:]
	fmt.Println(len(s8), cap(s8))

	// len(s) = 5
	// cap(s) = 10
	// len(s8) = 3
	// cap(s8) = 8 = cap(s) - i
	// s9 = s[i : j]
	s = make([]int, 5, 10)
	s9 := s[2:5]
	fmt.Println(len(s9), cap(s9))

	// clear
	clear(s)

	// 
	// s[0:len(s1)]
}
