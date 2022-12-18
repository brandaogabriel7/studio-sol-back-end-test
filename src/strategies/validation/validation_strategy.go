package validation

type ValidationStrategy interface {
	IsValid(password string, value int) bool
}