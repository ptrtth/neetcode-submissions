import ( "slices" )

func getSortedChars(s string) []byte {
	chars := []byte(s)
	slices.Sort(chars)
	return chars
}

func isAnagram(s string, t string) bool {
	return string(getSortedChars(s)) == string(getSortedChars(t))

}
