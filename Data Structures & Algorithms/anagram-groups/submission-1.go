import "slices"

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))

	for _, word := range strs {
		key := sortString(word)
		anagrams[key] = append(anagrams[key], word)
	}

	result := make([][]string, 0, len(anagrams))

	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}