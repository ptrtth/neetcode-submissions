func lengthOfLongestSubstring(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		seen := map[byte]struct{}{}

		for j := i; j < len(s); j++ {
			current := s[j]

			if _, exists := seen[current]; exists {
				break
			}

			seen[current] = struct{}{}
			result = max(result, len(seen))
		}
	}

	return result
}