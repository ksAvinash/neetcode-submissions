func maxArea(heights []int) int {
	top := 0
	for i, _ := range heights {
		for j, _ := range heights {
			area := min(heights[i], heights[j]) * int(math.Abs(float64(j-i)))
			// fmt.Println("area", heights[i], heights[j], area)
			if area > top {
				top = area
			}
		}
	}
	return top
}
