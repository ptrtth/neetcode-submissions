func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && matrix[i][len(matrix[i]) - 1] >= target {
			for _, value := range matrix[i] {
				if value == target {
					return true
				} else if value > target {
					break
				}
			}
		}
	}
	return false
}
