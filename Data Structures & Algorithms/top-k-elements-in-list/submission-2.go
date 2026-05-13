import "slices"

func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	freqs := map[int][]int{}

	for _, v := range nums {
		counts[v]++
		freqs[counts[v]] = append(freqs[counts[v]], v)
	}
	// fmt.Println(counts, freqs)

	i, res := len(freqs), []int{}
	for i > 0 {
		for _, q := range freqs[i] {
			if !slices.Contains(res, q) {
				res = append(res, q)
				if len(res) == k {
					return res
				}
			}
		}
		i--
	}
	// fmt.Println(res)
	return res
}