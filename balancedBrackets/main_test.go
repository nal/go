package main

import (
	"fmt"
	"testing"
)

func TestVerifyBrackets(t *testing.T) {
	type testcase struct {
		input string
		want  bool
	}

	tests := []testcase{
		{
			input: "(a[0]+b[2c[6]]) {24 + 53}",
			want:  true,
		},
		{
			input: "f(e(d))",
			want:  true,
		},
		{
			input: "[()]{}([])",
			want:  true,
		},
		{
			input: "((b)",
			want:  false,
		},
		{
			input: "(c]",
			want:  false,
		},
		{
			input: "{(a[])",
			want:  false,
		},
		{
			input: "([)]",
			want:  false,
		},
		{
			input: ")(",
			want:  false,
		},
		{
			input: "",
			want:  true,
		},
		{
			input: "{aa((([1][2][3])))}",
			want:  true,
		},
	}

	for _, test := range tests {
		want := test.want
		if got := verifyBrackets(test.input); got != want {
			t.Errorf("verifyBrackets() = %v, want %v, input = %s\n", got, want, test.input)
		} else {
			fmt.Printf("Passed test verifyBrackets() = %v, want %v, input = %s\n", got, want, test.input)
		}

	}

}
