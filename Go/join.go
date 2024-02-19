package piscine

func Join(strs []string, sep string) string {
	strV := ""

	for i := 0; i < len(strs); i++ {
		strV += strs[i]
		if i < len(strs)-1 {
			strV += sep
		}
	}
	return strV
}
