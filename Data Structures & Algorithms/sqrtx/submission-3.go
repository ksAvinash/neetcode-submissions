func mySqrt(x int) int {
	// brute force
	if x < 2 {
		return x
	}
	val := 0
	for i := range x {
		if i * i <= x {
			val = i
		} else {
			break
		}
	}
	return val
}
