package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0 && s[i] <= 31 {
			return false
		}
	}
	return true
}
