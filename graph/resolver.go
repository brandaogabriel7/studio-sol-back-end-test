package graph

import "github.com/brandaogabriel7/studio-sol-back-end-test/src/services/password_validation"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	PasswordValidationService password_validation.PasswordValidationService
}
