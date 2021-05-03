package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 1; i <= nb/2; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb < 0 {
		return false
	}
	for i := 2; i <= Sqrt(nb); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
