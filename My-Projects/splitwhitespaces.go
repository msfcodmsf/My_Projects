package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	a := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if a != "" {
				arr = append(arr, a)
				a = ""
			}
		} else {
			a += string(char)
		}
	}

	if a != "" {
		arr = append(arr, a)
	}

	return arr
}
