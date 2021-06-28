package lexorank

func isValid(s string) bool {
	for _, s := range []byte(s) {
		if !(s >= '0' && s <= '9') && !(s >= 'A' && s <= 'Z') && !(s >= 'a' && s <= 'z') {
			return false
		}
	}
	return true
}

func getChar(s string, i int, defaultChar byte) byte {
	if i >= len(s) {
		return defaultChar
	}
	return s[i]
}

func avg(prev, next byte) byte {
	return (prev + next) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
