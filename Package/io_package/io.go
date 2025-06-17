package main

import (
	"io"
	"os"
)

func main() {
	io.Copy(os.Stdout, os.Stdin)

	io.MultiReader()
	io.MultiWriter()

	// Incomplete
}