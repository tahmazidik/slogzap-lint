package helpers

import "unicode"

func FirstUnicodeLetter(s string) (rune, bool) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return r, true
		}
	}
	return 0, false
}

func ContainsCyrillic(s string) bool {
	for _, r := range s {
		if r >= 0x0400 && r <= 0x04FF {
			return true
		}
	}
	return false
}
