func convertToTitle(columnNumber int) string {
	res := []rune{}

	for columnNumber > 0 {
		// convert to base 25
		columnNumber--
		// take remainder
		rem := columnNumber % 26
		// convert rem to rune and append to list
		res = append(res, ('A' + rune(rem)))
		// go to the next number
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
