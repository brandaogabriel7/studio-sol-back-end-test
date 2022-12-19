package validation

import "regexp"

// Checks that the password contains at least the minimum number of special chars.
// The special chars are these: !@#$%^&*()-+\/{}[]
type MinSpecialCharsStrategy struct {
	RegexValidation
}

func NewMinSpecialCharsValidationStrategy() *MinSpecialCharsStrategy {
	const SPECIAL_CHARS_REGEXP = `[!@#$%^&*()\-+\\\/{}\[\]]`
	return &MinSpecialCharsStrategy{RegexValidation: RegexValidation{validationExpression: *regexp.MustCompile(SPECIAL_CHARS_REGEXP)}}
}