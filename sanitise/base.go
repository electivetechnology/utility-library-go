package sanitise

func SplitByCommaSpace(r rune) bool {
	return r == ',' || r == ' '
}

func SplitBySlash(r rune) bool {
	return r == '/'
}

func TrimByZero(r rune) bool {
	return r == '0'
}
