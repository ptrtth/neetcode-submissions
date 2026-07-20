import "slices"

func maxSlidingWindow(nums []int, k int) []int {
	results := []int{}
	for i := 0; i <= len(nums) - k; i++ {
		results = append(results, slices.Max(nums[i:i+k]))
	}
	return results
}
