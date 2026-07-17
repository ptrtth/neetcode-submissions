import "slices"

func trap(height []int) int {
	total := 0

	for i := 1; i < len(height)-1; i++ {
		leftMax := slices.Max(height[:i])
		rightMax := slices.Max(height[i+1:])

		if water := min(leftMax, rightMax) - height[i]; water > 0 {
			total += water
		}
	}

	return total
}