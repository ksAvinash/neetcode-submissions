func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var counts [26]int
		for _, ch := range s {
			counts[ch-'a']++
		}
		groups[counts] = append(groups[counts], s)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
