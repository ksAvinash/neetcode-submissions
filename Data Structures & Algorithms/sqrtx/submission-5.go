func mySqrt(x int) int {
	// brute force
	// if x < 2 {
	// 	return x
	// }
	// val := 0
	// for i := range x {
	// 	if i * i <= x {
	// 		val = i
	// 	} else {
	// 		break
	// 	}
	// }
	// return val


	if x < 2 {
		return x
	}
	left, right := 1, x
	val := 1
	for left <= right {
		mid := (left+right)/2

		res := mid * mid
		if res == x {
			return mid
		}
		if res < x {
			val = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return val
}
