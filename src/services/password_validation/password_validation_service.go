package password_validation

import (
	"github.com/brandaogabriel7/studio-sol-back-end-test/graph/model"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

type PasswordValidationService struct {
	strategies map[string]validation.ValidationStrategy
}

func NewPasswordValidationService(strategies map[string]validation.ValidationStrategy) *PasswordValidationService {
	return &PasswordValidationService{strategies: strategies}
}

func (pvs *PasswordValidationService) Validate(password string, rules []*model.Rule) model.Verify {
	verifyResponse := model.Verify{Verify: true, NoMatch: make([]string, 0)}

	for _, rule := range rules {
		if strategy, exists := pvs.strategies[rule.Rule]; exists {
			if !strategy.IsValid(password, rule.Value) {
				verifyResponse.NoMatch = append(verifyResponse.NoMatch, rule.Rule)
			}
		}
	}
	verifyResponse.Verify = len(verifyResponse.NoMatch) == 0

	return verifyResponse
}