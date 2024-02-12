package piscine

func ReverseMenuIndex(menu []string) []string {
	answer := make([]string, len(menu))
	j := 0
	for i := len(menu) - 1; i >= 0; i-- {
		answer[j] += menu[i]
		j++
	}
	return answer
}
