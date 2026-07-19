import "slices"

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)

	sortedS1 := []byte(s1)
	slices.Sort(sortedS1)

	for i := 0; i <= len(s2)-n; i++ {
		current := []byte(s2[i : i+n])
		slices.Sort(current)

		if slices.Equal(current, sortedS1) {
			return true
		}
	}

	return false
}