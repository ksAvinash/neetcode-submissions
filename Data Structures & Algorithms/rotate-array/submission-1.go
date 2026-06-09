func rotate(nums []int, k int) {
	l := len(nums)
	// for i := range k {
	// 	nums = append([]int{nums[l-1]}, nums[:l-1]...)
	// 	fmt.Println(i, nums)
	// }

	
	pos := map[int]int{}
	for i, v := range nums {
		pos[i] = v 
	}
	// fmt.Println(pos)

	k = k%l
	i, j := 0, k
	for i < l {
		if i < k {
			nums[i] = pos[l-j]
			j--
		} else {
			nums[i] = pos[i-k]
		}
		// fmt.Println(i, nums)
		i++
	}
}
