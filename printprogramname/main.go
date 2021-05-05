package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	t := []rune(string(os.Args[0]))
	for i := 2; i < len(t); i++ {
		z01.PrintRune(t[i])
	}
	z01.PrintRune('\n')
}