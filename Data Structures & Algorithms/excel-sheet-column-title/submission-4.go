func convertToTitle(columnNumber int) string {
	res := []rune{}

	for columnNumber > 0 {
		columnNumber--
		rem := columnNumber % 26
		res = append(res, ('A' + rune(rem)))
		columnNumber = columnNumber / 26
	}

	i, j := 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
