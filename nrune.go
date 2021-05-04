package piscine

func NRune(s string, n int) rune {
	t := []rune(s)
	if 0 < n && n <= len(s) {
		return t[n-1]
	}
	return 0
}
