import (
	"cmp"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	values := make([]int, 0, len(freq))

	for num := range freq {
		values = append(values, num)
	}

	slices.SortFunc(values, func(a, b int) int {
		if freq[a] != freq[b] {
			return cmp.Compare(freq[b], freq[a])
		}

		return cmp.Compare(a, b)
	})

	return values[:k]
}