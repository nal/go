package rot13

import "testing"

var tests = []struct {
	plain  string
	secret string
}{
	{
		plain:  "A",
		secret: "N",
	},
	{
		plain:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		secret: "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm",
	},
	{
		// Add some unicode characters, they should not be modified by algorithm
		plain:  "\u20acA",
		secret: "\u20acN",
	},
	{
		plain:  "Why did the chicken cross the road?\nGb trg gb gur bgure fvqr!",
		secret: "Jul qvq gur puvpxra pebff gur ebnq?\nTo get to the other side!",
	},
}

var implementations = []rot13{
	rot13ArrayShift{},
	rot13Switch{},
	rot13StringsMap{},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		for index := range implementations {
			r13 := implementations[index]
			got := r13.Encode(test.plain)
			if got != test.secret {
				t.Errorf("%T implementation got %q, wanted %q", r13, got, test.secret)
			}
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range tests {
		for index := range implementations {
			r13 := implementations[index]
			got := r13.Encode(test.secret)
			if got != test.plain {
				t.Errorf("%T implementation got %q, wanted %q", r13, got, test.plain)
			}
		}
	}
}

func BenchmarkEncodeRot13ArrayShift(b *testing.B) {
	r13 := rot13ArrayShift{}
	for n := 0; n < b.N; n++ {
		r13.Encode("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}
}

func BenchmarkEncodeRot13Switch(b *testing.B) {
	r13 := rot13Switch{}
	for n := 0; n < b.N; n++ {
		r13.Encode("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}
}

func BenchmarkEncodeRot13StringsMap(b *testing.B) {
	r13 := rot13StringsMap{}
	for n := 0; n < b.N; n++ {
		r13.Encode("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}
}
