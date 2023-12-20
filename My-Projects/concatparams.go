package piscine

func ConcatParams(args []string) string {
	yeni := ""
	for i, a := range args {
		yeni += string(a)
		if i < len(args)-1 {
			yeni += "\n"
		}
	}
	return yeni
}
