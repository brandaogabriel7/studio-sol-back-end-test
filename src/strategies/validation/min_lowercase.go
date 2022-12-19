package validation

import "regexp"

// Checks that the password has at least the minimum number of lower case characters.
type MinLowercaseValidationStrategy struct {
	RegexValidation
}

func NewMinLowercaseValidationStrategy() MinLowercaseValidationStrategy {
	const MIN_LOWERCASE_REGEXP string = "[a-z]"
	return MinLowercaseValidationStrategy{RegexValidation: RegexValidation{
		validationExpression: *regexp.MustCompile(MIN_LOWERCASE_REGEXP),
	}}
}