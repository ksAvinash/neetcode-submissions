func maxArea(heights []int) int {
	i, j := 0, len(heights)-1
	max := 0

	for i <= j {
		vol := min(heights[i], heights[j]) * (j-i)
		if vol > max {
			max = vol
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
