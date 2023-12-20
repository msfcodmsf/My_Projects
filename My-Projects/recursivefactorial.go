package piscine

func RecursiveFactorial(nb int) int {
	result := 1

	if nb == 0 || nb == 1 {
		return 1
	} else if nb > 1 && nb < 22 {
		result = nb * RecursiveFactorial(nb-1)
		return result
	} else {
		return 0
	}
}
