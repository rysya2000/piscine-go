package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var answer string = ""
	var str string = ""
	if n < 0 {
		n *= -1
		answer += "-"
	}
	p := n
	for p > 0 {
		str += ((p%10) + 48)
		p /= 10
	}
	for i := 0; i < len(str); i++ {
		str[i] = str[len(str)-i]
	}
	answer += str
	for i := 0; i < len(answer); i++ {
		z01.PrintRune(rune(answer[i]))
	}
}