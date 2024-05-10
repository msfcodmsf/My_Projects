package piscine

func IsNumeric(s string) bool {
	bool := false
	for _, i := range s {
		if i >= '0' && i <= '9' {
			bool = true
		} else {
			bool = false
			return bool
		}
	}
	return bool
}
