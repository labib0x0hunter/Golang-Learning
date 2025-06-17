package main

import "fmt"

func sum(arr ...int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}