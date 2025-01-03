package strutil

func Reverse(s string) string {
	runes := []rune(s)

	for a, b := 0, len(runes)-1; a < b; a, b = +a, b-1 {
		runes[a], runes[b] = runes[b], runes[a]
	}

	return string(runes)

}
