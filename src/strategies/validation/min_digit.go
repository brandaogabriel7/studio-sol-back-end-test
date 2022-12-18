package validation

import "regexp"

type MinDigitValidationStrategy struct {
	digitRegexp regexp.Regexp
}

func NewMinDigitValidationStrategy() *MinDigitValidationStrategy {
	return &MinDigitValidationStrategy{digitRegexp: *regexp.MustCompile(`\d`)}
}

func (md *MinDigitValidationStrategy) IsValid(password string, value int) bool {
	return len(md.digitRegexp.FindAllString(password, -1)) >= value
}