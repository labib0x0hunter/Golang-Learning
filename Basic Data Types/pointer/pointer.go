package main

func main() {

	var x int

	p := &x
	*p = 100 // dereferencing

	var q *int = p
	*q = 200
}