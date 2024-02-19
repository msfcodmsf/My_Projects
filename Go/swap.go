package piscine

func Swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
