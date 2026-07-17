import "slices"

func trap(height []int) int {

	n := len(height)

	if len(height) < 2 {
		return 0
	}

	result := 0

	for i := 1; i < n - 1; i++ {
		current := height[i]
		left := height[:i]
		right := height[i:]
		diff := min(slices.Max(left), slices.Max(right)) - current
		if diff > 0 {
			result = result + diff
		}
	}

	return result
}
