package piscine

func Compact(ptr *[]string) int {
	size := 0
	slice := *ptr

	for _, val := range slice {
		if val != "" {
			slice[size] = val
			size++
		}
	}

	*ptr = slice[:size]

	return size
}
