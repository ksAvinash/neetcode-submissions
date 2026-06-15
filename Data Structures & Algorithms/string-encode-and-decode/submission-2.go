type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "nil"
	}
	return strings.Join(strs, "@@")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "nil" {
		return []string{}
	}
	return strings.Split(encoded, "@@")
}
