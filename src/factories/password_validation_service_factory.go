package factories

import (
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/services/password_validation"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

// Creates the a PasswordValidation service with the following rule validation strategies:
// minSize, minSpecialChars, minDigit, noRepeted
func GetDefaultPasswordValidationService() password_validation.PasswordValidationService {
	validationStrategies := map[string]validation.ValidationStrategy{
		string(validation.MIN_SIZE): validation.MinSizeValidationStrategy{},
		string(validation.MIN_DIGIT): validation.NewMinDigitValidationStrategy(),
		string(validation.MIN_SPECIAL_CHARS): validation.NewMinSpecialCharsValidationStrategy(),
		string(validation.NO_REPETED): validation.NoRepetedStrategy{},
		string(validation.MIN_UPPERCASE): validation.NewMinUppercaseValidationStrategy(),
		string(validation.MIN_LOWERCASE): validation.NewMinLowercaseValidationStrategy(),
	}

	return *password_validation.NewPasswordValidationService(validationStrategies)
}