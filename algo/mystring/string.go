package mystring

func MatchingPairs(s string, t string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			// strings.
			println("s i ", string(t[i]))
		}
	}
	return 0
}
