package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinUppercase", func() {
	minUppercaseStrategy := validation.NewMinUppercaseValidationStrategy()

	Context("Check that the password follows minUppercase rule", func ()  {
		DescribeTable("When password contains more upper case characters than minUppercase value",
			func (password string, minUppercase int)  {
				isValid := minUppercaseStrategy.IsValid(password, minUppercase)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minUppercase 5", "SENHAA", 5),
			Entry("minUppercase 10", "SUPERSENHA123", 10),
			Entry("minUppercase 14", "123PARALELEPIPEDOSENHA", 14),
		)

		DescribeTable("When password contains as many upper case characters as minUppercase value",
			func (password string, minUppercase int)  {
				isValid := minUppercaseStrategy.IsValid(password, minUppercase)

				Expect(isValid).To(BeTrue())
			},
			Entry("minUppercase 5", "SENHA", 5),
			Entry("minUppercase 10", "SUPERSENHA123", 10),
			Entry("minUppercase 23", "DISFARCANDOASEVIDENCIAS", 23),
		)
	})

	Context("Check that the password does not follow minUppercase rule", func ()  {
		DescribeTable("When password has less upper case characters than minUppercase value",
		func (password string, minUppercase int)  {
			isValid := minUppercaseStrategy.IsValid(password, minUppercase)

			Expect(isValid).To(BeFalse())
		},
		Entry("minUppercase 5", "opaaa", 5),
		Entry("minUppercase 10", "superSenha321", 10),
		Entry("minUppercase 15", "SSSSuperS3nh@Forte", 15),
		)
	})
})
