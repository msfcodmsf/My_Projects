package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	cesce := true
	decresce := true

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			cesce = false
		}
	}

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			decresce = false
		}
	}
	return cesce || decresce
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
