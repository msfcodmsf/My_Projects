package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb == 0 || nb == 1 {
		return 1
	} else if nb > 1 && nb < 22 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	} else {
		return 0
	}
	return result
}
