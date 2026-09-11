func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, (m*n)-1

	for low <= high {
		mid := (low+high)/2
		val := matrix[int(mid/n)][mid%n]
		if val == target {
			return true
		} else if val > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}