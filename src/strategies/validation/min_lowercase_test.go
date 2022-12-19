package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinLowercase", func() {
	minLowercaseStrategy := validation.NewMinLowercaseValidationStrategy()

	Context("Check that the password follows minLowercase rule", func ()  {
		DescribeTable("When password contains more lower case characters than minLowercase value",
			func (password string, minLowercase int)  {
				isValid := minLowercaseStrategy.IsValid(password, minLowercase)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minLowercase 3", "senha123456", 3),
			Entry("minLowercase 10", "minhasenhaaaa12", 10),
			Entry("minLowercase 23", "eusouumasenhasuperforteesegura", 23),
		)

		DescribeTable("When password contains as many lower case characters as minLowercase value",
			func (password string, minLowercase int)  {
				isValid := minLowercaseStrategy.IsValid(password, minLowercase)

				Expect(isValid).To(BeTrue())
			},
			Entry("minLowercase 5", "senha12345", 5),
			Entry("minLowercase 10", "sseenhaaaa12345", 10),
			Entry("minLowercase 18", "eusouumasupersenha", 18),
		)
	})

	Context("Check that the password does not follow minLowercase rule", func ()  {
		DescribeTable("When password has less lower case characters than minLowercase value",
		func (password string, minLowercase int)  {
			isValid := minLowercaseStrategy.IsValid(password, minLowercase)

			Expect(isValid).To(BeFalse())
		},
		Entry("minLowercase 5", "opa", 5),
		Entry("minLowercase 10", "superSenha321", 10),
		Entry("minLowercase 30", "SuperS3nh@Forte", 30),
		)
	})
})
