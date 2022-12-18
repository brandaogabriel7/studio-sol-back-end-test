package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinDigit", func() {
	minDigitStrategy := validation.MinDigitValidationStrategy{}

	Describe("Check that the password follows minDigit rule", func ()  {
		DescribeTable("When password length is more than minDigit value",
			func (password string, minDigit int)  {
				isValid := minDigitStrategy.IsValid(password, minDigit)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minDigit 5", "password", 5),
			Entry("minDigit 10", "senhaaaa12345", 10),
			Entry("minDigit 13", "SuperS3nh@Forte12345", 13),
		)

		DescribeTable("When password length is equal to minDigit value",
			func (password string, minDigit int)  {
				isValid := minDigitStrategy.IsValid(password, minDigit)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minDigit 5", "opaaa", 5),
			Entry("minDigit 10", "superSenha", 10),
			Entry("minDigit 15", "SuperS3nh@Forte", 15),
		)
	})

	Describe("Check that the password does not follow minDigit rule", func ()  {
		DescribeTable("When password length is less than minDigit value",
		func (password string, minDigit int)  {
			isValid := minDigitStrategy.IsValid(password, minDigit)

			Expect(isValid).To(BeFalse())
		},
		Entry("minDigit 7", "opa", 7),
		Entry("minDigit 20", "senhaaaa12345", 20),
		Entry("minDigit 70", "SuperS3nh@Forte12345", 70),
		)
	})
})
