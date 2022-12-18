package validation

type MinSizeValidationStrategy struct {}

func (md *MinSizeValidationStrategy) IsValid(password string, value int) bool {
	return len(password) >= value
}