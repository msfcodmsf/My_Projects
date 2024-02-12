package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if max > min {
		answer := make([]int, len)
		for i := 0; i < len; i++ {
			answer[i] = min
			min++
		}
		return answer
	}
	return nil
}
