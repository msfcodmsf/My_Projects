package piscine

func IsLower(s string) bool {
	bool := false
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			bool = true
		} else {
			bool = false
			return bool
		}
	}
	return bool
}
