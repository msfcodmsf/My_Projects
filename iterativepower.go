package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}
	res := nb
	for i := 0; i < power-1; i++ {
		res = res * nb
	}
	return res
}
