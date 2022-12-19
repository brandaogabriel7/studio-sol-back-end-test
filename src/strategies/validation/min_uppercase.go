package validation

import "regexp"

type MinUppercaseValidationStrategy struct{
	RegexValidation
}

func NewMinUppercaseValidationStrategy() MinUppercaseValidationStrategy {
	const MIN_UPPERCASE_REGEXP string = "[A-Z]"
	return MinUppercaseValidationStrategy{RegexValidation: RegexValidation{
		validationExpression: *regexp.MustCompile(MIN_UPPERCASE_REGEXP),
	}}
}