package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}
	res := nb

	res = res * RecursivePower(nb, power-1)
	return res
}
