package piscine

func SortIntegerTable(table []int) {
	n := len(table)

	for i := 0; i < n/2; i++ {
		table[i], table[n-i-1] = table[n-i-1], table[i]
	}
}
