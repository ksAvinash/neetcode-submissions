func isValidSudoku(board [][]byte) bool {
	counts := make(map[string]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			val := string(board[i][j])
			if val == "." {
				continue
			}

			b := (i/3)*3 + j/3
			row := fmt.Sprintf("row-%d-%s", i, val)
			col := fmt.Sprintf("col-%d-%s", j, val)
			box := fmt.Sprintf("box-%d-%s", b, val)

			if counts[row] || counts[col] || counts[box] {
				return false
			}
			counts[row] = true
			counts[col] = true
			counts[box] = true
		}
	}
	return true
}
