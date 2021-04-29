package piscine

func StrRev(s string) string {
	rns := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		rns[i],rns[len(s)-1-i] = rns[len(s)-1-i], rns[i]
	}
	return string(rns)
}
