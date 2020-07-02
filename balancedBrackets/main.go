package main

func main() {
	/*
		Write a function that returns true if the brackets in a given string are balanced.
		The function must handle parens (), square brackets [], and curly braces {}.
		#### Instructions for Candidate:
		* Solve for the problem.
		* Add any additional tests/conditions that you feel are missing from the list of tests.
		* To run the tests, create a run config to execute class Solution and do not forget to enable assertions.
		* Depending on your IDE, you should see a "VM options" field, please add: -ea
		// (a[0]+b[2c[6]]) {24 + 53} -> true
		// f(e(d)) -> true
		// [()]{}([]) -> true
		// ((b) -> false
		// (c] -> false
		// {(a[]) -> false
		// ([)] -> false
		// )( -> false
		// -> true
	*/
}

func verifyBrackets(input string) (bracketsBalanced bool) {
	// Keep order of the brackets
	type bracketStruct struct {
		wasOpened bool
		counter   int
	}

	// Counterparts of brackets
	counterparts := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	// Map to store temorary results
	charBrackets := make(map[string]bracketStruct)

	// Store previous character to return early if counterparted bracket was open
	// before closing of other type bracket.
	var prevChar string

	for _, char := range input {
		stringChar := string(char)

		// Closing bracket
		if stringChar == ")" || stringChar == "]" || stringChar == "}" {
			value := charBrackets[counterparts[stringChar]]

			// Return early if we found closing bracket of the same type before opening
			if !value.wasOpened {
				return
			}

			// Look for opened bracket of different type before closing bracket
			// Eg: "([)". Fix for case at the last position of this string.
			for k, v := range counterparts {
				// Skip current bracket type
				if k == stringChar {
					continue
				}

				if prevChar == v {
					return
				}
			}

			value.counter--
			charBrackets[counterparts[stringChar]] = value

		} else if stringChar == "(" || stringChar == "[" || stringChar == "{" {
			// Opening bracket
			value := charBrackets[stringChar]
			value.wasOpened = true
			value.counter++
			charBrackets[stringChar] = value
		}

		prevChar = stringChar
	}

	for _, value := range charBrackets {
		if value.counter != 0 {
			return
		}
	}

	bracketsBalanced = true
	return
}
