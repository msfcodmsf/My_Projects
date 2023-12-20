package piscine

func UltimateDivMod(a *int, b *int) {
	var temp int
	temp = *a
	*a = *a / *b
	*b = temp % *b
}
