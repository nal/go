package rot13

import (
	"bytes"
)

// Implementation with array shift
type rot13ArrayShift struct{}

// Check type implementation
var _ rot13 = rot13ArrayShift{}

// Encode implements ROT13 encoding and decoding in one method `Encode()`.
// ROT13 algorithm description: https://en.wikipedia.org/wiki/ROT13.
// Encode converts plain text to "secret" text.
func (rot13 rot13ArrayShift) Encode(plain string) string {
	// Defines a value we use for shifting index to get new value.
	const shiftIndex = 13

	// Total alphabet capacity. 26 uppercase and 26 lowercase letters.
	const alphabetLen = 26 + 26

	alphabet := [alphabetLen]byte{
		// Uppercase
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		// Lowercase
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	dst := make([]rune, 0, len(plain))
	for _, chr := range plain {
		haveLatinChar := false
		subset := make([]byte, alphabetLen/2)

		// A-Z
		if chr > 64 && chr < 91 {
			haveLatinChar = true
			subset = alphabet[0 : alphabetLen/2]
		}

		// a-z
		if chr > 96 && chr < 123 {
			haveLatinChar = true
			subset = alphabet[alphabetLen/2:]
		}

		if haveLatinChar {
			currIndex := bytes.IndexByte(subset, byte(chr))
			if currIndex == -1 {
				panic("index not found for rune")
			} else {
				newIndex := currIndex + shiftIndex
				if newIndex > alphabetLen/2-1 {
					newIndex = newIndex - alphabetLen/2
				}
				dst = append(dst, rune(subset[newIndex]))
			}
		} else {
			dst = append(dst, chr)
		}
	}

	return string(dst)
}

// Decode converts "secret" text to plain text.
// In fact `Decode()` method calls `Encode()` method because of the nature of ROT13 algorithm.
func (rot13 rot13ArrayShift) Decode(secret string) string {
	return rot13.Encode(secret)
}
