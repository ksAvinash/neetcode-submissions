func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}


func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	}

	i, j, q := 0, k-1, k
	for q < len(arr) {
		if abs(arr[q]-x) < abs(arr[i]-x) {
			i++
			j++
		}
		q++
	}

	res := []int{}
	for i <= j {
		res = append(res, arr[i])
		i++
	}
	return res
}
