package main

import "fmt"

// len(s) == cap(s),
// In this situation append in s
// increases the capacity by 2 * len(s)
func main() {
	s := []int{1}

	fmt.Println(len(s), cap(s))
	s = append(s, 3)
	fmt.Println(len(s), cap(s))
	s = append(s, 4)
	fmt.Println(len(s), cap(s))
	s = append(s, 5)
	fmt.Println(len(s), cap(s))
	s = append(s, 6)
	fmt.Println(len(s), cap(s))
}

// s[:] referencing the slice of s which is extendable to s's capacity
// so changing on s[:] in range 0 to len(s) - 1, affects s (change the value)
// after append if slices length becomes  more than the capacity of s
// then s doesn't affected
// len(s) = 2
// cap(s) = 4
func Slice() {
	s := make([]int, 2, 4)
	Print2(s[:])
	fmt.Println("main", s)
}

// Affects
// because the memory of s[:] = c is the actual s
// s[:] = c length doesn't exceed capacity of s
// changes s[0] = 300
// len(c) = 3
// cap(s) = 4
func Print(c []int) {
	fmt.Println(c, len(c), cap(c))
	fmt.Println(c)
	c = append(c, 10)
	fmt.Println(c)

	c[0] = 3000
	fmt.Println(c)
}

// Affects
// because the memory of s[:] = c is the actual s
// s[:] = c length doesn't exceed capacity of s
// changes s[0] = 500
// len(c) = 4
// cap(s) = 4
func Print1(c []int) {
	c[0] = 3000
	c = append(c, 20)
	c = append(c, 30)
	c[0] = 5000
	fmt.Println(c)
}

// Doesn't Affects
// because the memory of s[:] = c is now changes as
// s[:] = c length exceed capacity of s
// len(c) = 6
// cap(s) = 4
func Print2(c []int) {
	c = append(c, 20)
	c = append(c, 30)
	c = append(c, 40)
	c = append(c, 50)
	c[0] = 5000
	fmt.Println(c)
}