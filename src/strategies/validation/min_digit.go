package validation

type MinDigitValidationStrategy struct {}

func (md *MinDigitValidationStrategy) IsValid(password string, value int) bool {
	return len(password) >= value
}