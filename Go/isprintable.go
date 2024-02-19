package piscine

func IsPrintable(s string) bool {
	bool := false
	for _, i := range s {
		if i >= 0 && i <= 31 {
			bool = false
			return bool
		} else {
			bool = true
		}
	}
	return bool
}
