import "maps"

func allCharsExists(chars map[byte]int) bool {
	for _, v := range chars {
		if v > 0 {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	target := map[byte]int{}

	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
    
	minLen := len(s)
	result := ""

	for i := 0; i < len(s); i++ {
		
		copied := maps.Clone(target)

		for j := i; j < len(s); j++ {
			if _, exists := copied[s[j]]; !exists {
				continue
			}

			copied[s[j]]--
			
			if !allCharsExists(copied) {
				continue
			}
			
			if minLen > j - i {
				result = string(s[i:j+1])
				minLen = len(result)
			}
		}
	}

	return result
}
