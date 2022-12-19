package factories

import (
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/services/password_validation"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

func GetDefaultPasswordValidationService() password_validation.PasswordValidationService {
	validationStrategies := map[string]validation.ValidationStrategy{
		string(validation.MIN_SIZE): validation.MinSizeValidationStrategy{},
		string(validation.MIN_DIGIT): validation.NewMinDigitValidationStrategy(),
		string(validation.MIN_SPECIAL_CHARS): validation.NewMinSpecialCharsValidationStrategy(),
		string(validation.NO_REPETED): validation.NoRepetedStrategy{},
	}

	return *password_validation.NewPasswordValidationService(validationStrategies)
}