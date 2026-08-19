type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, s := range strs {
		res += fmt.Sprintf("%d", len(s)) + "#" + s
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	i, res := 0, []string{}
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		slen, _ := strconv.Atoi(encoded[i:j])
		res = append(res, encoded[j+1 : j+1+slen])
		i = j+1+slen
	}
	return res
}
