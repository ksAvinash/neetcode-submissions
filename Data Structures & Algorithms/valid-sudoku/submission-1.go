func isValidSudoku(board [][]byte) bool {
	seen := map[string]int{}

	for r := range(9) {
		for c := range(9) {
			val := string(board[r][c])
			if val == "." {
				continue
			}
			rowKey := fmt.Sprintf("row-%d-val-%d", r, val)
			colKey := fmt.Sprintf("col-%d-val-%d", c, val)
			boxKey := fmt.Sprintf("box-%d-%d-val-%d", r/3, c/3, val)

			if seen[rowKey] > 0 || seen[colKey] > 0 || seen[boxKey] > 0 {
				return false
			}
			seen[rowKey] = 1
			seen[colKey] = 1
			seen[boxKey] = 1
 		}
	}
	

	return true
}
