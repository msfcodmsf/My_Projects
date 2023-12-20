package piscine

func Fibonacci(index int) int {
	res := 1
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	res = Fibonacci(index-2) + Fibonacci(index-1)
	return res
}
