func isValidSudoku(board [][]byte) bool {
	rows := map[string]map[string]int{}
	cols := map[string]map[string]int{}
	boxs := map[string]map[string]int{}

	var ex bool
	for r := range(9) {
		// fmt.Println("start", r, rows)
		for c := range(9) {
			val := string(board[r][c])
			if val == "." {
				continue
			}
			rk := fmt.Sprintf("row-%d", r)
			ck := fmt.Sprintf("col-%d", c)
			bk := fmt.Sprintf("box-%d-%d", r/3, c/3)

			_, rex := rows[rk][val]
			_, cex := cols[ck][val]
			_, bex := boxs[bk][val]
			if rex || cex || bex {
				return false
			}

			_, ex = rows[rk]; if !ex {
				rows[rk] = map[string]int{}
			}
			_, ex = cols[ck]; if !ex {
				cols[ck] = map[string]int{}
			}
			_, ex = boxs[bk]; if !ex {
				boxs[bk] = map[string]int{}
			}
			rows[rk][val] = 1
			cols[ck][val] = 1
			boxs[bk][val] = 1

 		}
		// fmt.Println("final", r, rows)
	}
	

	return true
}
