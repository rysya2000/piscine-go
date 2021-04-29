package piscine

import "github.com/01-edu/z01"

func IsNegative(a int) {
	if a >= 0 {
		z01.PrintRune("F\n")
	} else {
		z01.PrintRune("T\n")
	}
}
