import "maps"

func checkInclusion(s1 string, s2 string) bool {
	counts := make(map[rune]int)

	for _, s := range s1 {
		counts[s]++
	}

	for i, s := range s2 {
		if _, ex := counts[s]; ex {
			items, j := maps.Clone(counts), i

			for j < len(s2) {
				ch := rune(s2[j])
				if _, ex := items[ch]; ex {
					items[ch]--
					if items[ch] == 0 {
						delete(items, ch)
					}
				} else {
					break
				}
				j++
			}

			if len(items) == 0 {
				return true
			}
		}
	}
	return false
}
