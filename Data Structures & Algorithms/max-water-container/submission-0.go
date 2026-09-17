func maxArea(heights []int) int {
	res := 0
	i := 0
	j := len(heights) - 1
	for i < j {
		capicaty := (j - i) * min(heights[i], heights[j])
		if res < capicaty {
			res = capicaty
		}
		if heights[i] < heights[j] {
			i++
		} else if heights[i] > heights[j] {
			j--
		} else {
			i++
			j--
		}
	}
	return res
}
