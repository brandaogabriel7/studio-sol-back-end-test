package validation

import "regexp"

// Checks that the password contains at least the minimum number of special chars.
// The special chars are these: !@#$%^&*()-+\/{}[]
type MinSpecialCharsStrategy struct {
	specialCharsRegexp regexp.Regexp
}

func NewMinSpecialCharsValidationStrategy() *MinSpecialCharsStrategy {
	const SPECIAL_CHARS_REGEXP = `[!@#$%^&*()\-+\\\/{}\[\]]`
	return &MinSpecialCharsStrategy{specialCharsRegexp: *regexp.MustCompile(SPECIAL_CHARS_REGEXP)}
}

func (msc MinSpecialCharsStrategy) IsValid(password string, value int) bool {
	return len(msc.specialCharsRegexp.FindAllString(password, -1)) >= value
}