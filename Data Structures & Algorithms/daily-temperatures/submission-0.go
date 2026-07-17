func dailyTemperatures(temperatures []int) []int {
	n:= len(temperatures)
	result := make([]int, n, n)

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
