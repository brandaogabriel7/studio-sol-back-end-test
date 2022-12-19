package password_validation_test

import (
	"github.com/brandaogabriel7/studio-sol-back-end-test/graph/model"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/services/password_validation"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

// ValidationStrategy test double
type mockedValidationStrategy struct { mock.Mock }

func newMockedValidationStrategy() *mockedValidationStrategy { return &mockedValidationStrategy{} }

func (m mockedValidationStrategy) IsValid(password string, value int) bool {
	args := m.Called(password, value)
	return args.Bool(0)
}

// tests
var _ = Describe("PasswordValidationService", func() {
	const FIRST_RULE string = "firstRule"
	const SECOND_RULE string = "secondRule"
	const THIRD_RULE string = "thirdRule"

	firstRuleStrategy := mockedValidationStrategy{}
	secondRuleStrategy := mockedValidationStrategy{}
	thirdRuleStrategy := mockedValidationStrategy{}

	mockedValidationStrategies := map[string]*mockedValidationStrategy{
		FIRST_RULE:  &firstRuleStrategy,
		SECOND_RULE: &secondRuleStrategy,
		THIRD_RULE:  &thirdRuleStrategy,
	}

	validationStrategies := make(map[string]validation.ValidationStrategy)

	for key, value := range mockedValidationStrategies {
		validationStrategies[key] = value
	}

	pvs := password_validation.NewPasswordValidationService(validationStrategies)

	DescribeTable("Validate password when following rules have not passed",
		func (password string, rules []*model.Rule, noMatch []string)  {
			// overwrite stub for the noMatch strategies
			for key, mockedStrategy := range mockedValidationStrategies {
				passsedValidation := !utils.Contains(noMatch, key)
				mockedStrategy.On("IsValid", password, mock.AnythingOfType("int")).Return(passsedValidation)
			}

			verifyResponse := pvs.Validate(password, rules)

			isValid := len(noMatch) > 0

			Expect(verifyResponse.NoMatch).To(Equal(noMatch))
			Expect(verifyResponse.Verify).To(Equal(isValid))
		},
		Entry(
			"third",
			"opa",
			[]model.Rule{
				{Rule: FIRST_RULE, Value: 3},
				{Rule: SECOND_RULE, Value: 2},
				{Rule: THIRD_RULE, Value: 0},
			},
			[]string{THIRD_RULE},
		),
		Entry(
			"first, second",
			"senhaaa",
			[]model.Rule{
				{Rule: FIRST_RULE, Value: 5},
				{Rule: SECOND_RULE, Value: 7},
				{Rule: THIRD_RULE, Value: 2},
			},
			[]string{FIRST_RULE, SECOND_RULE},
		),
		Entry(
			"-",
			"senhaforte",
			[]model.Rule{
				{Rule: FIRST_RULE, Value: 0},
				{Rule: SECOND_RULE, Value: 7},
				{Rule: THIRD_RULE, Value: 2},
			},
			[]string{},
		),
	)
})
