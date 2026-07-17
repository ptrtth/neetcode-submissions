import "slices"

type Triplet struct {
	i, j, k int
}

func threeSum(nums []int) [][]int {
	results := make(map[Triplet]struct{})

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] != 0 {
					continue
				}

				values := []int{nums[i], nums[j], nums[k]}
				slices.Sort(values)

				triplet := Triplet{
					i: values[0],
					j: values[1],
					k: values[2],
				}

				results[triplet] = struct{}{}
			}
		}
	}

	output := make([][]int, 0, len(results))

	for triplet := range results {
		output = append(output, []int{
			triplet.i,
			triplet.j,
			triplet.k,
		})
	}

	return output
}