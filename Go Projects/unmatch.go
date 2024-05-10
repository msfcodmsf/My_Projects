package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		c := 0
		for _, el := range a {
			if el == res {
				c++
			}
		}
		if c == 1 || c%2 == 1 {
			return res
		}
	}
	return -1
}
