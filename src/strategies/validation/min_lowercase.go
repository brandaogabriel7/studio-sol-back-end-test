package validation

import "regexp"

type MinLowercaseValidationStrategy struct {
	RegexValidation
}

func NewMinLowercaseValidationStrategy() MinLowercaseValidationStrategy {
	const MIN_LOWERCASE_REGEXP string = "[a-z]"
	return MinLowercaseValidationStrategy{RegexValidation: RegexValidation{
		validationExpression: *regexp.MustCompile(MIN_LOWERCASE_REGEXP),
	}}
}