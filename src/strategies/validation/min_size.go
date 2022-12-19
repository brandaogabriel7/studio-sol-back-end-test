package validation

// Checks that the password is at least the minimum size.
type MinSizeValidationStrategy struct {}

func (md MinSizeValidationStrategy) IsValid(password string, value int) bool {
	return len(password) >= value
}