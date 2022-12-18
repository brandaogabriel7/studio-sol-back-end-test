package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("MinSize", func() {
	minSizeStrategy := validation.MinSizeValidationStrategy{}

	Describe("Check that the password follows minSize rule", func ()  {
		DescribeTable("When password length is more than minSize value",
			func (password string, minSize int)  {
				isValid := minSizeStrategy.IsValid(password, minSize)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minSize 5", "password", 5),
			Entry("minSize 10", "senhaaaa12345", 10),
			Entry("minSize 13", "SuperS3nh@Forte12345", 13),
		)

		DescribeTable("When password length is equal to minSize value",
			func (password string, minSize int)  {
				isValid := minSizeStrategy.IsValid(password, minSize)
	
				Expect(isValid).To(BeTrue())
			},
			Entry("minSize 5", "opaaa", 5),
			Entry("minSize 10", "superSenha", 10),
			Entry("minSize 15", "SuperS3nh@Forte", 15),
		)
	})

	Describe("Check that the password does not follow minSize rule", func ()  {
		DescribeTable("When password length is less than minSize value",
		func (password string, minSize int)  {
			isValid := minSizeStrategy.IsValid(password, minSize)

			Expect(isValid).To(BeFalse())
		},
		Entry("minSize 7", "opa", 7),
		Entry("minSize 20", "senhaaaa12345", 20),
		Entry("minSize 70", "SuperS3nh@Forte12345", 70),
		)
	})
})
