package piscine

func RockAndRoll(n int) string {
	var str string
	if n < 0 {
		str = "error: number is negative\n"
		return str
	}

	if n%2 == 0 && n%3 == 0 {
		str = "rock and roll"
	} else if n%2 == 0 {
		str = "rock"
	} else if n%3 == 0 {
		str = "roll"
	} else {
		str = "error: non divisible\n"
		return str
	}
	str += "\n"
	return str
}
