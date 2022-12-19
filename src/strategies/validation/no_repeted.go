package validation

import "strings"

// Checks that there are no consecutive repeated characters in the given password.
type NoRepetedStrategy struct {}

func (nr NoRepetedStrategy) IsValid(password string, _ int) bool {
	compressedPasswordSb := &strings.Builder{}
	var previous rune
	for _, r := range password {
		if r != previous {
			compressedPasswordSb.WriteRune(r)
			previous = r
		}
	}
	return compressedPasswordSb.String() == password
}