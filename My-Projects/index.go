package piscine

func Index(s string, toFind string) int {
	i := []rune(s)
	j := []rune(toFind)
	k := len(i)
	l := len(j)
	for a := 0; a < k-l; a++ {
		if toFind == s[a:a+l] {
			return (a)
		}
	}
	return -1
}
