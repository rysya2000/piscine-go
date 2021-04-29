package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
    var answer string = ""
    var str string = ""
    if n < 0 {
        z01.PrintRune('-')
    }
    for {
        var k int = (n % 10)
		if k < 0 {
			k *= -1
		}
        str += string(k + 48)
        n /= 10
    }
    rns := []rune(str)
    for i := 0; i < len(rns)/2; i++ {
        rns[i], rns[len(rns)-1-i] = rns[len(rns)-i-1], rns[i]
    }
    answer += string(rns)
    for i := 0; i < len(answer); i++ {
        z01.PrintRune(rune(answer[i]))
    }
}
