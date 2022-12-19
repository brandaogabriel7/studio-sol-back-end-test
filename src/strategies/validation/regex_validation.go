package validation

import "regexp"

// Checks that the password has at least the minimum number of whatever is specified in the validationExpression.
type RegexValidation struct {
	validationExpression regexp.Regexp
}

func (rv RegexValidation) IsValid(password string, value int) bool {
	return len(rv.validationExpression.FindAllString(password, -1)) >= value
}