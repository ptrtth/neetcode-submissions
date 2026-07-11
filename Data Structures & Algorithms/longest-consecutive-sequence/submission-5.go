import "slices"

func longestConsecutive(nums []int) int {

	n := len(nums)

	if n <= 1 {
		return n
	} 

	slices.Sort(nums)
	mx := 1
	cur := 1
	
	for i := 0; i < n-1; i++ {
		switch nums[i+1] - nums[i] {
		case 1:
			cur++
			mx = max(mx, cur)
		case 0:
			continue
		default:
			cur = 1
		}
	}

	return mx

}
