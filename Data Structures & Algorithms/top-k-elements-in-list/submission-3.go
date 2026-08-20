func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, 0)
	for _, n := range nums {
		freq[n]++
	}
	// fmt.Println(freq)

	counts := make([][]int, len(nums)+1, len(nums)+1)
	for k, v := range freq {
		counts[v] = append(counts[v], k)
	}
	// fmt.Println(counts)

	j, res := len(counts)-1, []int{}
	for j > 0 {
		if len(counts[j]) > 0 {
			res = append(res, counts[j]...)
		}
		if len(res) >= k {
			break
		}
		j--
	}

	return res

}
