func largestRectangleArea(heights []int) int {
	maxSize := 0

	for i := 0; i < len(heights); i++ {
		minHeight := math.MaxInt
		for j := i; j < len(heights); j++ {
			minHeight = min(minHeight, heights[j])
			maxSize = max(maxSize, minHeight * (j - i + 1))
		}
	}

	return maxSize
}
