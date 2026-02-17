package helpers

func FirstForbiddenSymbol(s string) (rune, bool) {
	for _, r := range s {
		if r == ' ' {
			continue
		}

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			continue
		}

		return r, true
	}
	return 0, false
}
