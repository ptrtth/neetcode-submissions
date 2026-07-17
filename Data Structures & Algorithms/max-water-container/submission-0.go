func maxArea(heights []int) int {
	result := 0

	for i := 0; i < len(heights) - 1; i++ {
		for j := i + 1; j < len(heights); j++ {
			current := min(heights[i], heights[j]) * (j - i)
			result = max(result, current)
		}
	}

	return result
}
