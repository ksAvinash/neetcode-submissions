func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	pos := map[int]int{}
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			pos[len(pos)] = nums1[i]
			i++
		} else {
			pos[len(pos)] = nums2[j]
			j++
		}
	}
	for i < m {
		pos[len(pos)] = nums1[i]
		i++
	}
	for j < n {
		pos[len(pos)] = nums2[j]
		j++
	}
	// fmt.Println(pos, i, j)
	for k, v := range pos {
		nums1[k] = v
	}
}
