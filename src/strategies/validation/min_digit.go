package validation

import "regexp"

// Checks that the password contains at least the minimum number of digits.
type MinDigitValidationStrategy struct {
	digitRegexp regexp.Regexp
}

func NewMinDigitValidationStrategy() *MinDigitValidationStrategy {
	const DIGIT_REGEXP string = `\d`
	return &MinDigitValidationStrategy{digitRegexp: *regexp.MustCompile(DIGIT_REGEXP)}
}

func (md MinDigitValidationStrategy) IsValid(password string, value int) bool {
	return len(md.digitRegexp.FindAllString(password, -1)) >= value
}