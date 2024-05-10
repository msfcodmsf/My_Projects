package piscine

func StringToIntSlice(str string) []int {
	//if len(str) == 0 {
	//	return nil
	//}
	//res := make([]int,0, len(str))
	var res []int

	for _, c := range str {
		res = append(res, int(c))
	}
	return res
}
