package validation

import "strings"

type NoRepetedStrategy struct {}

func (nr *NoRepetedStrategy) IsValid(password string, _ int) bool {
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