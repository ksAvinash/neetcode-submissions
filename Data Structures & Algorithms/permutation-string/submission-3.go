func checkInclusion(s1 string, s2 string) bool {
	s1counts := map[string]int{}
	for _, v := range s1 {
		s1counts[string(v)]++
	}

	for i := 0; i < len(s2); i++ {
		_, ex := s1counts[string(s2[i])]
		if ex {
			counts := make(map[string]int, len(s1counts))
			for k, v := range s1counts { counts[k] = v }

			j, valid := i, true
			for j < i+len(s1) {
				if j >= len(s2) {
					valid = false
					break
				}
				v, ex := counts[string(s2[j])]
				if v == 0 || !ex {
					valid = false
					break
				}
				counts[string(s2[j])]--
				j++
			}
			if valid {
				return true
			}
		}
	}

	return false
}
