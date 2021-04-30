package piscine

func Atoi(s string) int {
	n := 0
	for index, ch := range []byte(s) {
		if index != 0 && (48 > ch || ch > 57) {
			return 0
		}
		if (48 > ch || ch > 57) && ch != '+' && ch != '-' {
			return 0
		}
		if 48 <= ch && ch <= 57 {
			ch -= 48
			n = n*10 + int(ch)
		}
	}
	if len(s) > 0 && s[0] == '-'{
		n = -n
	}
	return int(n)
}
