package piscine

func CountIf(f func(string) bool, tab []string) int {
	l := 0
	for _, s := range tab {
		if f(s) == true {
			l++
		}
	}
	return l
}
