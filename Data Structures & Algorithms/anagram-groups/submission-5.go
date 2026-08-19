import "slices"

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, s := range strs {
		sorted := sortString(s)
		_, ex := anagrams[sorted]
		if ex {
			anagrams[sorted] = append(anagrams[sorted], s)
		} else {
			anagrams[sorted] = []string{s}
		}
	}
	res := [][]string{}
	for _, v := range anagrams {
		res = append(res, v)
	}
	return res
}
