func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, (m*n)-1

	for low <= high {
		mid := int((low+high)/2)
		i, j := int(mid/n), mid%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
