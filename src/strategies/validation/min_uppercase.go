package validation

import "regexp"

// Checks that password has at least the minimum number of lower case characters.
type MinUppercaseValidationStrategy struct{
	RegexValidation
}

func NewMinUppercaseValidationStrategy() MinUppercaseValidationStrategy {
	const MIN_UPPERCASE_REGEXP string = "[A-Z]"
	return MinUppercaseValidationStrategy{RegexValidation: RegexValidation{
		validationExpression: *regexp.MustCompile(MIN_UPPERCASE_REGEXP),
	}}
}