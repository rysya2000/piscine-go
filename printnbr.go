package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var answer string = ""
	var str string = ""
	var p int = n
	if n < 0 {
		n *= -1
		answer += "-"
	}
	for p > 0 {
		var k int = (p % 10) + 48
		str += string(k)
		p /= 10
	}
	for i := 0; i < len(str) / 2; i++ {
		str[i], str[len - i - 1] = str[len(str)-i - 1], str[i]
	}
	answer += str
	for i := 0; i < len(answer); i++ {
		z01.PrintRune(rune(answer[i]))
	}
}
