package piscine

import "github.com/01-edu/z01"

func toStr(n int) string{
	if n == 0 {
        return "0"
    }
	ans := ""
	var x int
	for n > 0 {
        x = n % 10
        n /= 10
        ans += string(rune(x + 48))
    }
	return ans
}

func PrintNbrInOrder(n int) {
	res := ""
	table := []byte(toStr(n))
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	for _, i := range table {
		res += string(i)
	}
	for i := 0; i < len(res); i++ {
		z01.PrintRune(rune(res[i]))
	}
}
dsadas