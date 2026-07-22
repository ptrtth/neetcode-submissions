func findMin(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	} else if nums[n - 1] > nums[0] {
		return nums[0]
	}

	result := 0

	for i := 0; i < n - 1; i++ {
		if nums[i + 1] < nums[i] {
			result = nums[i + 1]
			break
		}
	}
	
	return result
}
