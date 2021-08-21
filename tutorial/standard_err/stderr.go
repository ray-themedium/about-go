package main

import (
	"io"
	"os"
)

func main() {
	str := ""
	args := os.Args
	if len(args) == 1 {
		str = "Please give one argument!"
	} else {
		str = args[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")
}
