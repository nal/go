package rot13

import (
	"strings"
)

// Implementation with strings.Map function and switch/case
type rot13StringsMap struct{}

// Check type implementation
var _ rot13 = rot13StringsMap{}

// Encode implements ROT13 encoding and decoding in one method `Encode()`.
// ROT13 algorithm description: https://en.wikipedia.org/wiki/ROT13.
// Encode converts plain text to "secret" text.
func (rot13 rot13StringsMap) Encode(src string) string {
	r13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	return strings.Map(r13, src)
}

// Decode converts "secret" text to plain text.
// In fact `Decode()` method calls `Encode()` method because of the nature of ROT13 algorithm.
func (rot13 rot13StringsMap) Decode(src string) string {
	return rot13.Encode(src)
}
