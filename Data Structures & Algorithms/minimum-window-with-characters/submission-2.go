import "maps"

func minWindow(s string, t string) string {
	target := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}

	bestStart := 0
	bestLen := len(s) + 1

	for left := 0; left < len(s); left++ {
		counts := maps.Clone(target)
		missing := len(t)

		for right := left; right < len(s); right++ {
			char := s[right]

			count, needed := counts[char]
			if needed {
				if count > 0 {
					missing--
				}

				counts[char]--
			}

			if missing == 0 {
				windowLen := right - left + 1

				if windowLen < bestLen {
					bestStart = left
					bestLen = windowLen
				}

				break
			}
		}
	}

	if bestLen == len(s)+1 {
		return ""
	}

	return s[bestStart : bestStart+bestLen]
}