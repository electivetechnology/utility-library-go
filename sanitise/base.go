package sanitise

import "github.com/electivetechnology/utility-library-go/logger"

var log logger.Logging

func SplitByCommaSpace(r rune) bool {
	return r == ',' || r == ' '
}

func SplitBySlash(r rune) bool {
	return r == '/'
}

func TrimByZero(r rune) bool {
	return r == '0'
}
