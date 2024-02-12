package piscine

func FindNextPrime(nb int) int {
	nextprime := nb - 1
	i := nb + 1
	for i > nb {
		nextprime++
		if IsPrime2(nextprime) {
			return nextprime
		}
		i++
	}
	return nextprime
}

func IsPrime2(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
