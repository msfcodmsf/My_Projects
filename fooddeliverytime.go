package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	}
	if order == "chips" {
		return 10
	}
	if order == "nuggets" {
		return 12
	}
	return 404
}
