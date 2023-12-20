package piscine

func IsUpper(s string) bool {
	bool := false
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			bool = true
		} else {
			bool = false
			return bool
		}
	}
	return bool
}
