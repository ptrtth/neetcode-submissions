func minEatingSpeed(piles []int, h int) int {
	rate := 1

	for {
		hours := 0

		for _, pile := range piles {
			hours += (pile + rate - 1) / rate

			if hours > h {
				break
			}
		}

		if hours <= h {
			return rate
		}

		rate++
	}
}