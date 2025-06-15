package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ifElse() {

	n := 100
	if n+n >= 200 {
		//
	} else if n >= 200 {
		//
	} else {
		//
	}

	if num, err := strconv.Atoi("324"); err != nil {
		fmt.Println("number: ", num)
	}
}

func switchCase() {
	cmd := "/"

	switch cmd {
	case "+":
		//
	case "-":
		//
	case "/":
		//
		fallthrough
	case "%":
		//
	default:
		//
	}

	r := rune('1')
	// r = rune('Ⅻ')
	switch {
	case unicode.IsLetter(r):
		fmt.Println("Rune is a letter")
	case unicode.IsDigit(r):
		fmt.Println("Rune is a digit")
		fallthrough
	case unicode.IsNumber(r):
		fmt.Println("Rune is a number")
	}
}

func main() {

	switchCase()

}
