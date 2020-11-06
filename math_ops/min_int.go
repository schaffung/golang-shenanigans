package min_int

func minInt(vals ...int) int {
	minVal := vals[0]
	for _, val := range vals[1:] {
		if minVal > val {
			minVal = val
		}
	}
	return minVal
}
