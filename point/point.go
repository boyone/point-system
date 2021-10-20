package point

func CalculatePoint(price float64) int {
	point := int(price) / 100
	return point
}
