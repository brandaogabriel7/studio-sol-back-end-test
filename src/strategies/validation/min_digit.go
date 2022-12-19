package validation

import "regexp"

// Checks that the password contains at least the minimum number of digits.
type MinDigitValidationStrategy struct {
	RegexValidation
}

func NewMinDigitValidationStrategy() *MinDigitValidationStrategy {
	const DIGIT_REGEXP string = `\d`
	return &MinDigitValidationStrategy{RegexValidation: RegexValidation{
		validationExpression: *regexp.MustCompile(DIGIT_REGEXP),
	}}
}