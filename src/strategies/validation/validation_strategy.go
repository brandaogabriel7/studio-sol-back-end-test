package validation

// ValidationStrategy defines a strategy to validate passwords based on a rule.
type ValidationStrategy interface {
	IsValid(password string, value int) bool
}