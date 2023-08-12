package rot13

// Implementation with switch/case
type rot13Switch struct{}

// Check type implementation
var _ rot13 = rot13Switch{}

// Encode implements ROT13 encoding and decoding in one method `Encode()`.
// ROT13 algorithm description: https://en.wikipedia.org/wiki/ROT13.
// Encode converts plain text to "secret" text.
func (rot13 rot13Switch) Encode(plain string) string {
	dst := make([]rune, 0, len(plain))
	for _, chr := range plain {
		if (chr > 64 && chr < 91) || // A-Z
			(chr > 96 && chr < 123) { // a-z
			switch chr {
			// Uppercase
			case 'A':
				dst = append(dst, 'N')
			case 'B':
				dst = append(dst, 'O')
			case 'C':
				dst = append(dst, 'P')
			case 'D':
				dst = append(dst, 'Q')
			case 'E':
				dst = append(dst, 'R')
			case 'F':
				dst = append(dst, 'S')
			case 'G':
				dst = append(dst, 'T')
			case 'H':
				dst = append(dst, 'U')
			case 'I':
				dst = append(dst, 'V')
			case 'J':
				dst = append(dst, 'W')
			case 'K':
				dst = append(dst, 'X')
			case 'L':
				dst = append(dst, 'Y')
			case 'M':
				dst = append(dst, 'Z')
			case 'N':
				dst = append(dst, 'A')
			case 'O':
				dst = append(dst, 'B')
			case 'P':
				dst = append(dst, 'C')
			case 'Q':
				dst = append(dst, 'D')
			case 'R':
				dst = append(dst, 'E')
			case 'S':
				dst = append(dst, 'F')
			case 'T':
				dst = append(dst, 'G')
			case 'U':
				dst = append(dst, 'H')
			case 'V':
				dst = append(dst, 'I')
			case 'W':
				dst = append(dst, 'J')
			case 'X':
				dst = append(dst, 'K')
			case 'Y':
				dst = append(dst, 'L')
			case 'Z':
				dst = append(dst, 'M')

			// Lowercase
			case 'a':
				dst = append(dst, 'n')
			case 'b':
				dst = append(dst, 'o')
			case 'c':
				dst = append(dst, 'p')
			case 'd':
				dst = append(dst, 'q')
			case 'e':
				dst = append(dst, 'r')
			case 'f':
				dst = append(dst, 's')
			case 'g':
				dst = append(dst, 't')
			case 'h':
				dst = append(dst, 'u')
			case 'i':
				dst = append(dst, 'v')
			case 'j':
				dst = append(dst, 'w')
			case 'k':
				dst = append(dst, 'x')
			case 'l':
				dst = append(dst, 'y')
			case 'm':
				dst = append(dst, 'z')
			case 'n':
				dst = append(dst, 'a')
			case 'o':
				dst = append(dst, 'b')
			case 'p':
				dst = append(dst, 'c')
			case 'q':
				dst = append(dst, 'd')
			case 'r':
				dst = append(dst, 'e')
			case 's':
				dst = append(dst, 'f')
			case 't':
				dst = append(dst, 'g')
			case 'u':
				dst = append(dst, 'h')
			case 'v':
				dst = append(dst, 'i')
			case 'w':
				dst = append(dst, 'j')
			case 'x':
				dst = append(dst, 'k')
			case 'y':
				dst = append(dst, 'l')
			case 'z':
				dst = append(dst, 'm')
			}
		} else {
			dst = append(dst, chr)
		}
	}

	return string(dst)
}

// Decode converts "secret" text to plain text.
// In fact `Decode()` method calls `Encode()` method because of the nature of ROT13 algorithm.
func (rot13 rot13Switch) Decode(secret string) string {
	return rot13.Encode(secret)
}
