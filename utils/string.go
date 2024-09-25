package utils

import "unicode"

func firstLetter(str string) rune {
	for _, c := range str {
		if unicode.IsLetter(c) {
			return c
		}
	}
	return '.'
}

func firstNumber(str string) string {
	for i, c := range str {
		if unicode.IsLetter(c) {
			return str[0 : i-1]
		}
	}
	return ""
}
