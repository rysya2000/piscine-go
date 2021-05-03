package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb > 0 && nb < 21 {
		return nb * IterativeFactorial(nb-1)
	}
	return 0
}
