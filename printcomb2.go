package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if i < 10 {
				z01.PrintRune(rune('0'))
				z01.PrintRune(rune(i + 48))
				z01.PrintRune(rune(' '))
			}
			if j < 10 {
				z01.PrintRune(rune('0'))
				z01.PrintRune(rune(j + 48))
			}
			if i > 9 {
				z01.PrintRune(rune((i / 10) + 48))
				z01.PrintRune(rune((i % 10) + 48))
				z01.PrintRune(rune(' '))
			}
			if j > 9 {
				z01.PrintRune(rune((j / 10) + 48))
				z01.PrintRune(rune((j % 10) + 48))
			}
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
