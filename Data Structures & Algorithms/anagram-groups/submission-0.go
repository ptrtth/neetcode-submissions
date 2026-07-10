import "slices"

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {

	anagrams := map[string][]string{}

	for _, s := range strs {
		id := sortString(s)
		anagrams[id] = append(anagrams[id], s)
	}

	result := make([][]string, len(anagrams))

	i := 0
	for _, v := range anagrams {
		result[i] = v
		i++
	}
	
	return result

}
