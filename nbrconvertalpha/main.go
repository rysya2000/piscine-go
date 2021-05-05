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

func IsValid(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args
	b := false
	upper := false
	var x int = 1
	if args[1] == "--upper" {
		upper = true
		x++
	}
	for i := x; i < len(args); i++ {
		if IsValid(args[i]) == false {
			b = true
			break
		}
	}
	if b == true {
		for i := x; i < len(args); i++ {
			z01.PrintRune(rune(' '))
		}
 	}
	if b == false {
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
	}
	z01.PrintRune(rune('\n'))
}