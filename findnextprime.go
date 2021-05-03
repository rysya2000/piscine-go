package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	for {
		b := 1
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				b = 0
			}
		}
		if b == 1 {
			return nb
		}
		nb += 1
	}
}
