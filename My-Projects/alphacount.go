package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			counter++
		}
	}
	return counter
}
