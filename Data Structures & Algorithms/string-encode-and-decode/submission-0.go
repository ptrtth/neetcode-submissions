type Solution struct{}

const ( 
	eos = "<eos>"
	empty = "<empty>"
)

func (s *Solution) Encode(strs []string) string {
	var result []string

	for _, s := range strs {
		if s == "" {
			result = append(result, empty)
		} else {
			result = append(result, s)
		}
		result = append(result, eos)
	}

	return strings.Join(result, "")

}

func (s *Solution) Decode(encoded string) []string {
	var result []string

	for _, s := range strings.Split(encoded, eos) {
		if s == empty {
			result = append(result, "")
		} else if s != "" {
			result = append(result, s)
		}
	}

	return result
}
