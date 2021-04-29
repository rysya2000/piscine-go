package piscine

func BasicAtoi(s string) int {
	rns := []rune(s)
	var ans int = 0
	k := 1
	for i := 0; i < len(s); i++ {
		k *= 10
	}
	for i := 0; i < len(s); i++ {
		ans += (int(rns[i]) - '0') * k
		k /= 10
	}
	return ans / 10
}
