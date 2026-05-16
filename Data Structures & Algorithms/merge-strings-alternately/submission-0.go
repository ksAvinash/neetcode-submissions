func mergeAlternately(word1 string, word2 string) string {
	i, j, ie, je := 0, 0, len(word1), len(word2)

	res := ""
	for i < ie && j < je {
		res += string(word1[i])
		res += string(word2[j])
		i++
		j++
	}
	if i < ie {
		res += word1[i:]
	}
	if j < je {
		res += word2[j:]
	}

	return res
}
