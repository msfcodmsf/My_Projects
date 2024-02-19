package piscine

func Capitalize(s string) string {
	result := []rune(s)
	CapitalizeNext := true

	for i, char := range result {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if CapitalizeNext {
				result[i] = ToUpperCase(char)
			} else {
				result[i] = ToLowerCase(char)
			}
			CapitalizeNext = false

		} else {
			CapitalizeNext = true
		}
	}
	return string(result)
}

func ToUpperCase(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func ToLowerCase(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}
