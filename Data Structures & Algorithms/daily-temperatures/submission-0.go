func dailyTemperatures(temperatures []int) []int {
	res := []int{}

	for i, _ := range temperatures {
		j := i + 1
		days := 0
		for j < len(temperatures) {
			if temperatures[j] > temperatures[i] {
				days = j - i
				break
			}
			j++
		}
		res = append(res, days)
	}
	return res
}
