package piscine

func Sqrt(nb int) int {
	sqrt := 1
	res := 1
	if nb == 0 || nb < 0 {
		return 0
	}

	for i := 1; i <= nb; i++ {

		sqrt = i * i
		if nb == sqrt {
			res = i
			return res
		}

	}
	return 0
}
