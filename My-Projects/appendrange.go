package piscine

func AppendRange(min, max int) []int {
	var answer []int
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	if min >= max {
		answer = append(answer)
	}
	return answer
}
