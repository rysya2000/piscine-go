package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		b := 1
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] && i+j < len(s) {
				b = 0
				break
			}
		}
		if b == 1 {
			return i
		}
	}
	return -1
}
