package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toInt(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= 48
		n = n*10 + int(ch)
	}
	return int(n)
}

func main() {
	args := os.Args
	upper := false
	var x int = 1
	if args[1] == "--upper" {
		upper = true
		x++
	}
	y := 97
	if upper == true {
		y -= 32
	}
	for i := x; i < len(args); i++ {
		if toInt(args[i]) <= 26 && toInt(args[i]) > 0 {
			z01.PrintRune(rune(toInt(args[i]) + y - 1))
		} else {
			z01.PrintRune(rune(' '))
		}
	}
	z01.PrintRune(rune('\n'))
}
