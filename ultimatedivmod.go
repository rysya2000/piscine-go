package piscine

func UltimateDivMod(a *int, b *int) {
	k := *a
	*a = *a / *b
	*b = k % *b
}
