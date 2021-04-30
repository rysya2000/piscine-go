package piscine

func BasicAtoi2(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		if 48 > ch || ch > 57 {
			return 0
		}
		ch -= 48
		n = n*10 + int(ch)
	}
	return int(n)
}
