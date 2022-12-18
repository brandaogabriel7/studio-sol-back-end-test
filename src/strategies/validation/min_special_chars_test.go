package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinSpecialChars", func() {
	minSpecialCharsStrategy := validation.NewMinSpecialCharsValidationStrategy()

	Describe("Check that the password follows minSpecialChars rule", func ()  {
		DescribeTable("When password contains more special chars than minSpecialChars value",
			func (password string, minSpecialChars int)  {
				isValid := minSpecialCharsStrategy.IsValid(password, minSpecialChars)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minSpecialChars 4", "senha!@%[]@!", 4),
			Entry("minSpecialChars 7", "senhaaaa[{!@}]^[]", 7),
			Entry("minSpecialChars 17", `SuperS3nh@Forte!@#$%^&*()-+\/{}[]`, 17),
		)

		DescribeTable("When password contains as many special chars as minSpecialChars value",
			func (password string, minSpecialChars int)  {
				isValid := minSpecialCharsStrategy.IsValid(password, minSpecialChars)

				Expect(isValid).To(BeTrue())
			},
			Entry("minSpecialChars 4", "senha@%[]", 4),
			Entry("minSpecialChars 7", "senhaaaa[!@}]^[", 7),
			Entry("minSpecialChars 19", `SuperS3nh@Forte!@#$%^&*()-+\/{}[]`, 19),
		)
	})

	Describe("Check that the password does not follow minSpecialChars rule", func ()  {
		DescribeTable("When password has less special chars than minSpecialChars value",
		func (password string, minSpecialChars int)  {
			isValid := minSpecialCharsStrategy.IsValid(password, minSpecialChars)

			Expect(isValid).To(BeFalse())
		},
		Entry("minSpecialChars 5", "opaaa", 5),
		Entry("minSpecialChars 10", "superSenha321-+^}", 10),
		Entry("minSpecialChars 15", "SuperS3nh@Forte", 15),
		)
	})
})