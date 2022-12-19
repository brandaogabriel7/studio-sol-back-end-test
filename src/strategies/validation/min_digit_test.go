package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinDigit", func() {
	minDigitStrategy := validation.NewMinDigitValidationStrategy()

	Context("Check that the password follows minDigit rule", func ()  {
		DescribeTable("When password contains more digits than minDigit value",
			func (password string, minDigit int)  {
				isValid := minDigitStrategy.IsValid(password, minDigit)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minDigit 5", "senha123456", 5),
			Entry("minDigit 10", "senhaaaa1234567891011", 10),
			Entry("minDigit 23", "1234567891011SuperS3nh@Forte123451234567891011", 23),
		)

		DescribeTable("When password contains as many digits as minDigit value",
			func (password string, minDigit int)  {
				isValid := minDigitStrategy.IsValid(password, minDigit)

				Expect(isValid).To(BeTrue())
			},
			Entry("minDigit 5", "senha12345", 5),
			Entry("minDigit 10", "senhaaaa1234567890", 10),
			Entry("minDigit 30", "1234567891011SuperS3nh@Forte123451234567891011", 30),
		)
	})

	Context("Check that the password does not follow minDigit rule", func ()  {
		DescribeTable("When password has less digits than minDigit value",
		func (password string, minDigit int)  {
			isValid := minDigitStrategy.IsValid(password, minDigit)

			Expect(isValid).To(BeFalse())
		},
		Entry("minDigit 5", "opaaa", 5),
		Entry("minDigit 10", "superSenha321", 10),
		Entry("minDigit 15", "SuperS3nh@Forte", 15),
		)
	})
})
