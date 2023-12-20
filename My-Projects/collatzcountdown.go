package piscine

func CollatzCountdown(start int) int {
	a := 0
	for start >= 1 {

		if start%2 == 0 {
			start = start / 2
			a++
		} else if start%2 == 1 {
			start = 3*start + 1
			a++
		}
		if start == 1 {
			return a
		}
	}
	return -1
}
