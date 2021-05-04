package piscine

func ToUpper(s string) string {
	t := []rune(s)
	for i := 0; i < len(s); i++ {
		if t[i] >= 97 && t[i] <= 122 {
			t[i] = t[i] - 32
		}
	}
	return string(t)
}
