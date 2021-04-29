package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 58; i++ {
		for j := i + 1; j < 58; j++ {
			for k := j + 1; k < 58; k++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
}
