package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadByte()
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()

	// Incomplete
}