package piscine

func Abort(a, b, c, d, e int) int {
	var item []int
	item = append(item, a, b, c, d, e)
	medyan := insertionsort(item)
	return medyan[2]
}

func insertionsort(items []int) []int {
	n := len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}
