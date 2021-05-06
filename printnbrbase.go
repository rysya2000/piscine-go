package piscine

func IsValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var a [1111]int
	for i := 0; i < len(s); i++ {
		a[s[i]]++
		if a[s[i]] > 1 {
			return false
		}
		if s[i] == '-' || s[i] == '+' {
			return false
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if IsValid(base) == false {
		z01.PrintRune(rune('N'))
		z01.PrintRune(rune('V'))
		return
	}
	ans := ""
	x := len(base)
	y := nbr % x
	if y < 0 {
		y = -y
	}
	ans += string(base[y])
	nbr /= x
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	for nbr > 0 {
		ans += string(base[nbr%x])
		nbr /= x
	}
	for i := len(ans) - 1; i >= 0; i-- {
		z01.PrintRune(rune(ans[i]))
	}
}
