package piscine

func IsAlpha(s string) bool {
	bool := false
	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '0' && i <= '9' {
			bool = true
		} else {
			bool = false
			return bool
		}
	}
	return bool
}
