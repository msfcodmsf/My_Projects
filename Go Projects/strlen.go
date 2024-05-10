package piscine

func StrLen(s string) int {
	l := 0
	a := []rune(s)
	for p := range a {
		l = p + 1
	}
	return l
}
