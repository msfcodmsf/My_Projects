package piscine

func IsPrime(nb int) bool {
	bool := false
	if nb == 2 {
		bool = true
		return bool
	} else if nb < 0 || nb == 1 {
		return bool
	}
	for i := 2; i < nb; i++ {
		if nb%i != 0 {
			bool = true
		} else {
			bool = false
			return bool
		}
	}
	return bool
}
